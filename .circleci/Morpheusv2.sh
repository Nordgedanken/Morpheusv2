#!/bin/bash
appname=`basename $0 | sed s,\.sh$,,`

export LD_LIBRARY_PATH=/usr/local/lib
export QT_PLUGIN_PATH=/usr/local/Morpheusv2/plugins
export QML_IMPORT_PATH=/usr/local/Morpheusv2/qml
export QML2_IMPORT_PATH=/usr/local/Morpheusv2/qml
/usr/bin/$appname "$@"