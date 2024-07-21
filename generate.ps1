$DOFUS_SCRIPTS_SOURCE_FOLDER="../dofus-decompiled-sources/14-06-2024/scripts"
$OUTPUT_DIRECTORY = "./pkg/decoder"

if (!(Test-Path -Path $OUTPUT_DIRECTORY -PathType Container)) {
    mkdir -p $OUTPUT_DIRECTORY
}

..\dofus-protocol-builder\dofus-protocol-builder.exe $DOFUS_SCRIPTS_SOURCE_FOLDER $OUTPUT_DIRECTORY

$DOFUS_DATA_FOLDER="c:\Users\batbo\AppData\Local\Ankama\Dofus\data\"
$OUTPUT_DIRECTORY = "./pkg/types"

go run .\cmd\types\types.go $DOFUS_DATA_FOLDER $OUTPUT_DIRECTORY
