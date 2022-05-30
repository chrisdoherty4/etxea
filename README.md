# Etxea ('house' in Basque)

A go-plugin proof of concept. The core set of packages under `pkg` house the components for building a plugin. Gela ('room' in Baseque) is a plugin that leverages the capabilities.

## Experiments

#### CLI flags

Can we provide pluggable CLI flags. A generic solution that can be bound to any arbitrary flag set would be useful so we can bind to multiple sub commands.

#### Cluster bootstrap

"Cluster bootstrap" is probably too specific. We might need to model stages in some capacity that leverage a provider but is there a general "bind at stage X" capability we could build? This may be hard dependent on inputs as they may differ from 1 stage vs another. Do we break this down per-plugin definition?

#### What happens when a plugin isn't configured?

Hashicorp's plugin mechanism lets you list plugins and 'dispense' a particular one. Is this an avenue to providing the different stages that we perhaps bind within?

On trying to use an unconfigured plugin the system simply returns an error.

```
Error: rpc error: code = Unimplemented desc = unknown service etxea.v1.BindingService
```

We could test for the `Unimplemented` code and look at the description (considering it comes from gRPC itself) to see if it matches a service. If the full service isn't implemented then we can ignore that stage and move on to the next making plugins opt in for authors.

#### Discovery

If the plugin framework is defined in xavier, can plugins leverage that and be discoverable? This seems to be how Terraform (and probably other HashiCorp products) work and makes sense given the plugin host needs to define expected behavior, not the plugin.

## Design considerations

- Inputs from 1 stage to the next
- Initialization and desctruction similar to a testing framework. Do we need per stage and 1 wrapper?
- Stage ordering.