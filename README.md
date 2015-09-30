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






