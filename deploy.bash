#!/bin/bash

directories=(
	"sp2/web/"
	"homepage/wserve/"
	"Uploads/"
)


commands=(
	"./server"
	"./homeweb"
	"./upload"

)

for ((i=0; i<${#directories[@]}; i++)); do
  (
    cd "${directories[i]}" && ${commands[i]} &
    PID=$!
    echo "Started command '${commands[i]}' in ${directories[i]} with PID $PID"
  )
done

sleep 1
