#!/bin/bash
./obsidian_vault-toc "$1"
obsidian-export "$1" "$2"
