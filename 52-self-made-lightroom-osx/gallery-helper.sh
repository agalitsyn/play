#!/usr/bin/env bash

function do_bulk_resize() {
    local size=${1:?"usage: doresize <size>"}
    mogrify -monitor -resize "$size" -gravity center -format jpg *.jpg
}

function do_thumbnail() {
    local src=${1:?"usage: do_thumbnail <src> <size>"}
    local size=${2:-"600"}
    local dst="${src%/*}"/thumb_"${src##*/}"
    convert -thumbnail "$size" -gravity center -format jpg "$src" "$dst"
    echo $dst
    echo
}
export -f do_thumbnail

function do_template() {
    local dir=$(basename "$(pwd)")
    local width=
    local height=

    echo '- var product = []'
    for file in $(find . -type f -name "IMG*" -printf '%P\n'); do
        width=$(identify -ping -format "%w" $file)
        height=$(identify -ping -format "%h" $file)
        echo "- product.push({imagePath: 'media/${dir}/${file}', imageProps: '${width}x${height}'})"
    done
    echo
}
export -f do_template

function do_bulk_template() {
    local folder=${1:?"usage: do_bulk_template <folder>"}

    find "$folder" -name *.jpg -printf '%h\n' |
        sort |
        uniq |
        xargs -I{} -0 -L1 cd $1 && do_template && cd -
}

function thumbnailize_summer() {
    summer="IMG_5132.jpg IMG_5365.jpg IMG_5200.jpg IMG_5043.jpg IMG_5752.jpg IMG_4995.jpg IMG_5190.jpg IMG_5065.jpg IMG_5109.jpg IMG_5257.jpg IMG_5641.jpg IMG_0282.jpg IMG_0318.jpg IMG_0081.jpg IMG_0374.jpg IMG_0171.jpg IMG_0114.jpg IMG_0198.jpg IMG_8861.jpg IMG_0423.jpg IMG_0433.jpg IMG_9055.jpg IMG_0796.jpg IMG_0906.jpg IMG_0821.jpg IMG_0806.jpg IMG_0688.jpg IMG_0786.jpg IMG_0505.jpg IMG_0529.jpg IMG_0477.jpg IMG_0489.jpg IMG_0539.jpg IMG_0561.jpg IMG_0468.jpg IMG_0659.jpg IMG_0594.jpg IMG_0973.jpg IMG_1044.jpg"
    find_and_thumbnail "$summer"
}

function thumbnailize_winter() {
    winter="IMG_7070.jpg IMG_7071.jpg IMG_6949.jpg IMG_6810.jpg IMG_6393.jpg IMG_6149.jpg IMG_5963.jpg IMG_5825.jpg IMG_6447.jpg IMG_6474.jpg IMG_6960.jpg IMG_6903.jpg IMG_6855.jpg IMG_6890.jpg IMG_6923.jpg IMG_6282.jpg IMG_6941.jpg IMG_5993.jpg"
    find_and_thumbnail "$winter"
}

function thumbnailize_demi() {
    demi="IMG_6378.jpg IMG_7750.jpg IMG_8340.jpg IMG_7704.jpg IMG_7735.jpg IMG_8456.jpg IMG_7685.jpg IMG_8231.jpg IMG_8573.jpg IMG_8688.jpg IMG_8795.jpg IMG_8122.jpg IMG_7989.jpg IMG_8888.jpg IMG_8974.jpg IMG_7913.jpg IMG_8920.jpg IMG_7665.jpg IMG_9646.jpg"
    find_and_thumbnail "$demi"
}

function thumbnailize_all() {
    thumbnailize_summer
    thumbnailize_winter
    thumbnailize_demi
}

function find_and_thumbnail() {
    local list=(${1:?"usage: find_and_thumbnail <list>"})

    for file in "${list[@]}"; do
        find . -name "$file" -printf '%P\n' | xargs -I{} -0 -L1 sh -c "echo \"==> \$0\"; do_thumbnail \$0"
    done
}
