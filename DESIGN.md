# Design

A discussion/raw notes on possible designs for plugin architecture.

```go
func main() {
	// 1. Perform initial flag parsing to retrieve configuration file given we're 'declarative'.
    // 2. Interpret the configuration to figure out what provider need to use.
	// 3. Launch the provider.
	// 4. Ask the provider to define any additional flags.
    // 5. Re-parse the arguments because the provider flags weren't previously defined. This is 
    //    necessary because we don't know what provider to use until we've performed a basic 
    //    parsing of the configuration file. Not all cmds need this. Some cmds like genererate
    //    cluster config get the provider from the CLI input.

	config, err := config.Parse(configPath)
    if err != nil {
        // handle
    }

    // Probably provide CLI configurability for paths withs some sensible default.
    // Default may be based on the program binary path.
	availableProvider := providers.Discover(paths)
	providerLauncher, err := availableProvider.Retrieve(config.ProviderPluginName())
	if err != nil {
		// handle
	}

	provider := providerLauncher.Launch()
    
    if provider.HasFlagService() {
        flags := provider.FlagService()
        flags.Populate(set)
    }
}
```