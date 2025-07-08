#!/usr/bin/php
<?php

if ( $argc != 3 ){
	echo "This program reads a open api json specification and fix invalid characters used in fragment.\n";
	echo "This program expects 2 arguments:\n";
	echo "1º name of input json file\n";
	echo "2º name of output json file\n";
	echo "\n\n\n";
	exit(1);
}

# CLI arguments
$argH = [];
$argH['input_file'] = $argv[1];
$argH['output_file'] = $argv[2];

$json = file_get_contents( $argH['input_file'] );
$assoc = TRUE;
$obj = json_decode( $json, $assoc );
if ( is_null($obj) ){
	echo "ERROR decoding json.\n";
	exit(2);
}

# iterate over component keys
$keyA = array_keys( $obj['components']['schemas'] );
$keyC = count($keyA);

echo "Found $keyC keys.\n";
$searchA = [];
$replaceA = [];
$re = "/[^a-zA-Z0-9\-._~!$&'()*+,;=:@\/?]/"; # test for invalid characters in fragment
for($i=0; $i<$keyC; $i++){
	$key = $keyA[$i];
	#echo $key,"\n";
	$r = preg_match( $re, $key );
	if ( $r === false ) {
		echo "Invalid regular expression [$re]";
		exit(1);
	}
	if ( $r == 1 ){
		#echo $key,"\n";
		$replacement = preg_replace( $re, '_', $key );
		$searchA[] = $key;
		$replaceA[] = $replacement;
	}
}
$cc = count($replaceA);
echo "Number of keys replaced: $cc\n";

# until here we have the list of replacements
# need to be applied to the original json
$json_mod = str_replace( $searchA, $replaceA, $json );

file_put_contents( $argH['output_file'], $json_mod );

