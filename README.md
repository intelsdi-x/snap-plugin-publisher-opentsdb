DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.
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


# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**



[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-publisher-opentsdb.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-opentsdb)

# Snap publisher plugin - OpenTSDB

Snap OpenTSDB plugin written in Go supports publishing ingested time series data points into OpenTSDB.
It's used in the [Snap framework](http://github.com:intelsdi-x/snap).


1. [Getting Started](#getting-started)
    * [System Requirements](#system-requirements)
    * [Installation](#installation)
    * [Configuration and Usage](#configuration-and-usage)
    * [Install and run OpenTSDB](#install-and-run-opentsdb)
2. [Documentation](#documentation)
    * [Examples](#examples)
    * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License and Authors](#license-and-authors)
6. [Thank You](#thank-you)


## Getting Started
To get started, you'll need Snap and OpenTSDB running to receive and aggregate sampling data points.

### System Requirements
* [golang 1.6+](https://golang.org/dl/) (needed only for building)
* [snap](https://github.com/intelsdi-x/snap)

### Operating systems
All OSs currently supported by Snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download opentsdb plugin binary:
You can get the pre-built binaries for your OS and architecture at plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb/releases) page.

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
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

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

Example of running [psutil collector plugin](https://github.com/intelsdi-x/snap-plugin-collector-psutil) and publishing data to OpenTSDB.

Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `sudo snapteld -l 1 -t 0 &`

Download and load Snap plugins (paths to binary files for Linux/amd64):
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-opentsdb/latest/linux/x86_64/snap-plugin-publisher-opentsdb
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-psutil/latest/linux/x86_64/snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-opentsdb
$ snaptel plugin load snap-plugin-collector-psutil
```

Create a [task manifest](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md) (see [exemplary tasks](examples/tasks/)),
for example `psutil-opentsdb.json` with following content:
```json
{
  "version": 1,
  "schedule": {
    "type": "simple",
    "interval": "10s"
  },
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/psutil/load/load1": {},
        "/intel/psutil/load/load15": {},
        "/intel/psutil/load/load5": {},
        "/intel/psutil/vm/available": {},
        "/intel/psutil/vm/free": {},
        "/intel/psutil/vm/used": {}
      },
      "publish": [
        {
          "plugin_name": "opentsdb",
          "config": {
            "host": "127.0.0.1",
            "port": 4242
          }
        }
      ]
    }
  }
}
```
Create a task:
```
$ snaptel task create -t psutil-opentsdb.json
```

Watch created task:
```
$ snaptel task watch <task_id>
```

To stop previously created task:
```
$ snaptel task stop <task_id>
```

#### Limitations
This plugin only supports the "host" tag for now.

#### Example Data
Here is an example of what psutil namespaces look like in tsdb.

| namespace | datatype | tag |
|-----------------------|:-------------------:|--------------:|
| intel.psutil.load.load1 | float64 | host |
| intel.psutil.load.load15 | float64 | host |
| intel.psutil.vm.used | float64 | host |
| intel.psutil.vm.free | float64 | host |

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-opentsdb/issues).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

## Thank You
Your contribution is incredibly important to us.






