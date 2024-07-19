$DOFUS_SCRIPTS_SOURCE_FOLDER="../dofus-decompiled-sources/14-06-2024/scripts"
$OUTPUT_DIRECTORY = "./pkg/decoder"

if (!(Test-Path -Path $OUTPUT_DIRECTORY -PathType Container)) {
    mkdir -p $OUTPUT_DIRECTORY
}

..\dofus-protocol-builder\dofus-protocol-builder.exe $DOFUS_SCRIPTS_SOURCE_FOLDER $OUTPUT_DIRECTORY
