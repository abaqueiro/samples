#!/usr/bin/php
<?php

if ( $argc != 4 ){
	echo "This program reads a open api json specification and add a server to it.\n";
	echo "This program expects 3 arguments:\n";
	echo "1º name of json file\n";
	echo "2º api url\n";
	echo "3º name of output file\n";
	echo "\n\n\n";
	exit(1);
}

# CLI arguments
$argH = [];
$argH['input_file'] = $argv[1];
$argH['api_url'] = $argv[2];
$argH['output_file'] = $argv[3];

$json = file_get_contents( $argH['input_file'] );
$assoc = TRUE;
$obj = json_decode( $json, $assoc );
if ( is_null($obj) ){
	echo "ERROR decoding json.\n";
	exit(2);
}

if ( ! isset($obj['servers']) ){
	$obj['servers'] = [];
	echo "[ INFO ] server property not found, adding\n";
}

$o2 = [];
$o2['url'] = $argH['api_url'];
$obj['servers'][] = $o2;

$json_mod = json_encode($obj);

echo "[ INFO ] writing modified json to {$argH['output_file']}\n";
file_put_contents( $argH['output_file'], $json_mod);

