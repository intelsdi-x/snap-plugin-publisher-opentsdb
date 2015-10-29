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

# pulse-plugin-publisher-opentsdb
Pulse Publisher Plugin to OpenTSDB

## Description
	This plugin publishes data into OpenTSDB for Pulse compliant collectors.

## Dependencies
	It requires project Pulse: https://github.com/intelsdi-x/pulse.

## Configuration
1. Set PULSE_PATH envoriment variable for running an example.
2. Set ULSE_OPENTSDB_HOST for integration test.
3. Run the example from the root directory. E.g. ./examples/run-opentsdb-psutil.sh <instance-name>

## Limitations
	1. Plugin only supports "host" tag for now and it's auto obtained from the running OS.

## Details
    Dot delimitered namespace. E.g.
	-----------------------------------------------------------------------
	| Namespace          | Datatype           | Tag
	-----------------------------------------------------------------------
	| psutil.load.1      | float64            |  hostname
	-----------------------------------------------------------------------

## Change log
	first PR 2015-9-21

## Licensing
	TBD

## Credits and acknowledgements

	Thanks for Joel Cooklin's help throughout the plugin development.






