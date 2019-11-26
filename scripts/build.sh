#!/bin/bash
set -e
echo "Generating template map from template files"
go generate ${PWD}/scripts/generate.go
echo "All files geneated"