#!/bin/bash
./obsidian-vault-toc "$1"
obsidian-export "$1" "$2"
