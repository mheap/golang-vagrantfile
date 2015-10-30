# vagrantfile

<img src="https://travis-ci.org/mheap/golang-vagrantfile.svg" />

### Why does this project exist?

For one of the projects I'm working on, we need to be able to generate
Vagrantfiles programatically. After a quick search nothing showed up on Github
so this project was born.

### Usage:

```golang
	v := &VagrantFile{
		Box:            "ubuntu/trusty64",
		BoxCheckUpdate: false,
		ForwardedPorts: []ForwardedPort{
			ForwardedPort{
				Guest: 80,
				Host:  8080,
			},

			ForwardedPort{
				Guest: 1234,
				Host:  5678,
			},
		},
	}

    output, err := v.Render()
```
