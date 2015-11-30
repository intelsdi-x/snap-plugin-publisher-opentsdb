<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->
[![Build Status](https://travis-ci.com/intelsdi-x/snap-plugin-publisher-opentsdb.svg?token=HoxHq3yqBGpySzRd5XUm&branch=master)](https://travis-ci.com/intelsdi-x/snap-plugin-publisher-opentsdb)

snap OpenTSDB plugin written in Go supports publishing ingested time series data points into OpenTSDB.
It's used in the [snap framework](http://github.com:intelsdi-x/snap).


1. [Getting Started](#getting-started)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3.  [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License and Authors](#license-and-authors)
6. [Thank You](#thank-you)


## Getting Started
To get started, you'll need snap and OpenTSDB running to receive and aggregate sampling data points.

### System Requirements
* [golang 1.4+](https://golang.org/dl/)
* [snap](https://github.com/intelsdi-x/snap)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download opentsdb plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-opentsdb.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

### Install and run OpenTSDB
You can install and run an OpenTSDB instance quickly using ​*the*​ Docker setup command below.
```
docker run -d --name opentsdb_opentsdb_1 -p 4242:4242 opower/opentsdb
```

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [snap opentsdb integration test](https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb/blob/master/opentsdb/opentsdb_integration_test.go)
* [snap opentsdb unit test](https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb/blob/master/opentsdb/opentsdb_test.go)

### Examples
Example running psutil plugin, passthru processor, and writing data into an opentsdb database.

Documentation for snap collector psutil plugin can be found [here](https://github.com/intelsdi-x/snap-plugin-collector-psutil)

In one terminal window, open the snap daemon :
```
$ snapd -t 0 -l 1
```
The option "-l 1" it is for setting the debugging log level and "-t 0" is for disabling plugin signing.

In another terminal window:

Load collector and processor plugins
```
$ snapctl plugin load $SNAP_PSUTIL_PLUGIN/build/rootfs/snap-plugin-collector-psutil
$ snapctl plugin load build/rootfs/plugin/snap-processor-passthru
```
Load opentsdb plugin
```
$ snapctl plugin load $SNAP_OPENTSDB_PLUGIN/build/rootfs/snap-plugin-publisher-opentsdb
```

See available metrics for your system
```
$ snapctl metric list
```

Create a task file. For example, sample-psutil-task.json:   
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/psutil/load/load1": {},
                "/intel/psutil/load/load5": {},
                "/intel/psutil/load/load15": {},
                "/intel/psutil/vm/available": {},
                "/intel/psutil/vm/free": {},
                "/intel/psutil/vm/used": {}
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "plugin_version": 1,
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "opentsdb",
                            "plugin_version": 1,
                            "config": {
                                "host": "OPENTSDB_IP",
                                "port": 4242                    
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```
Create task:
```
$ snapctl task create -t sample-psutil-task.json
```
### Limitations
This plugin only supports the "host" tag for now.

## Example Data
Here is an example of what psutil namespaces look like in tsdb.

| namespace | datatype | tag |
|-----------------------|:-------------------:|--------------:|
| intel.psutil.load.load1 | float64 | host |
| intel.psutil.load.load15 | float64 | host |
| intel.psutil.vm.used | float64 | host |
| intel.psutil.vm.free | float64 | host |

### Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb/issues).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

## Thank You
Your contribution is incredibly important to us.






