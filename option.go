package ship

type Option struct {
	DryRun bool `long:"dry-run" short:"d" destination:"Dry run. Only print invoice."`
}
