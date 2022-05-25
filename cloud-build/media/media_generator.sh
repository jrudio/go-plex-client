#! /bin/bash
#
# this script will create mock movies and tv shows for PMS to pick up and add to the library
input_file=$1
target_directory=$2

if [ -z "$input_file" ] && [[ $input_file != *.mp4 ]] && [[ $input_file != *.mkv ]]; then
  echo "Usage: $0 <input_file>.mp4 <target_directory>"
  exit 1
fi

if [ -z "$target_directory" ]; then
  target_directory="."
fi

echo "generating fake media using '$input_file'..."

# start with a few and possibly expand later
mkdir -p "$target_directory/movies"
mkdir -p "$target_directory/tv"

cp $input_file $target_directory/movies/Interstellar\ \(2014\).mp4
cp $input_file $target_directory/movies/Interstellar\ \(2014\).mp4

mkdir -p "$target_directory/tv/Dave Chappelle's Show/Season 2"

cp $input_file "$target_directory/tv/Dave Chappelle's Show/Season 2/S02E05.mp4"

echo "finished generating media in '$target_directory'"