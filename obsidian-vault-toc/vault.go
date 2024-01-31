package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/md"
	"github.com/gomarkdown/markdown/parser"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Vault struct {
	RootPath    string
	SummaryFile string
	Directories []string
	Files       []*VaultFile
}

type VaultFile struct {
	Parent string
	Path   string
	Token  string
}

func find(root, patt string) ([]string, error) {
	var walk []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(patt, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			walk = append(walk, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return walk, err
}

func dirs(root string) ([]string, error) {
	var walk []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			walk = append(walk, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return walk, err
}

func contains(fi string, li []string) bool {
	for _, item := range li {
		if fi == item {
			return true
		}
	}
	return false
}

func allSummarized(files []*VaultFile, li []string) bool {
	for _, f := range files {
		if !contains(f.Token, li) {
			return false
		}
	}
	return true
}

func titleize(s string) string {
	t := strings.ReplaceAll(s, "_", " ")
	rs := []rune(t)
	inWord := false
	for i, r := range rs {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if !inWord {
				rs[i] = unicode.ToTitle(r)
			}
			inWord = true
		} else {
			inWord = false
		}
	}
	return string(rs)
}

func getTitle(file *VaultFile) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	noExt := strings.TrimSuffix(file.Token, filepath.Ext(file.Token))

	b, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return titleize(noExt)
	}
	var title string
	f := p.Parse(b)
	ast.WalkFunc(f, func(node ast.Node, entering bool) ast.WalkStatus {
		if h, ok := node.(*ast.Heading); ok && entering {
			t := h.GetChildren()
			if len(t) == 1 {
				title = string(t[0].AsLeaf().Literal)
				return ast.Terminate
			}
		}
		return ast.GoToNext
	})
	if title == "" {
		title = titleize(noExt)
		s := fmt.Sprintf("# %s\n%s", title, string(b))
		err := ioutil.WriteFile(file.Path, []byte(s), 0644)
		if err != nil {
			panic(err)
		}
	}
	return title
}

func writeDirectoryTOC(d, root, path string, files []*VaultFile) error {
	mdText := fmt.Sprintf("# %s", d)
	for _, f := range files {
		mdText = fmt.Sprintf("%s\n- [%s](%s)", mdText, getTitle(f), f.Token)
	}
	err := ioutil.WriteFile(filepath.Join(root, path), []byte(mdText), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadVault(path string) (*Vault, error) {
	vault := &Vault{RootPath: path}
	var summaryPath = filepath.Join(path, "SUMMARY.md")
	if _, err := os.Stat(summaryPath); !os.IsNotExist(err) {
		vault.SummaryFile = summaryPath
	}
	directories, err := dirs(path)
	if err != nil {
		return vault, err
	}
	vault.Directories = directories
	paths, err := find(path, "*.md")
	if err != nil {
		return vault, err
	}
	for _, path := range paths {
		vf := &VaultFile{}
		vf.Path = path
		vf.Token = filepath.Base(path)
		var dir = filepath.Dir(path)
		if filepath.Base(dir) != filepath.Base(vault.RootPath) {
			vf.Parent = dir
		}
		vault.Files = append(vault.Files, vf)
	}
	return vault, nil
}

func SummarizeVault(vault *Vault) error {
	if vault.SummaryFile == "" {
		return fmt.Errorf("need summary file (SUMMARY.md) for vault at path '%s'", vault.RootPath)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	body, err := ioutil.ReadFile(vault.SummaryFile)
	if err != nil {
		panic(err)
	}
	summary := p.Parse(body)
	var summarized []string
	var list ast.Node
	ast.WalkFunc(summary, func(node ast.Node, entering bool) ast.WalkStatus {
		if _, ok := node.(*ast.List); ok && entering {
			if list == nil {
				list = node
			}
		}
		if link, ok := node.(*ast.Link); ok && entering {
			url := link.Destination
			summarized = append(summarized, filepath.Base(string(url)))
		}
		return ast.GoToNext
	})
	buckets := make(map[string][]*VaultFile)
	for _, file := range vault.Files {
		parent := filepath.Base(file.Parent)
		if len(buckets[parent]) == 0 {
			buckets[parent] = make([]*VaultFile, 0)
		}
		buckets[parent] = append(buckets[parent], file)
	}
	for d, files := range buckets {
		var anchor ast.Node
		if d != "" && d != "." {
			dtoc := fmt.Sprintf("%s.md", d)
			if err := writeDirectoryTOC(d, vault.RootPath, dtoc, files); err != nil {
				return err
			}
		}

		if len(files) == 0 || allSummarized(files, summarized) {
			continue
		}
		if d != "" && d != "." {
			dtoc := fmt.Sprintf("%s.md", d)
			node := &ast.ListItem{BulletChar: '-'}
			para := &ast.Paragraph{}
			link := &ast.Link{Destination: []byte(dtoc)}
			sublist := &ast.List{Tight: true}
			node.SetChildren([]ast.Node{para, sublist})
			para.SetChildren([]ast.Node{&ast.Text{}, link, &ast.Text{}})
			link.SetChildren([]ast.Node{&ast.Text{Leaf: ast.Leaf{Literal: []byte(d)}}})
			ast.AppendChild(list, node)
			anchor = sublist
		} else {
			continue
		}
		for _, file := range files {
			if !contains(file.Token, summarized) {
				node := &ast.ListItem{BulletChar: '-'}
				para := &ast.Paragraph{}
				link := &ast.Link{Destination: []byte(file.Token)}
				node.SetChildren([]ast.Node{para})
				para.SetChildren([]ast.Node{&ast.Text{}, link, &ast.Text{}})
				link.SetChildren([]ast.Node{&ast.Text{Leaf: ast.Leaf{Literal: []byte(getTitle(file))}}})
				ast.AppendChild(anchor, node)
			}
		}
	}
	ast.Print(os.Stdout, summary)
	mdRenderer := md.NewRenderer()
	mdText := markdown.Render(summary, mdRenderer)
	err = ioutil.WriteFile(vault.SummaryFile, mdText, 0644)
	if err != nil {
		return err
	}
	return nil
}
