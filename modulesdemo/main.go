package main

import "goBasic/modulesdemo/exports"

func main() {
	export := exports.New("Jaumne", "Singer")
	export.ExportMethod()
}
