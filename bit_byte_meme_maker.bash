#for word in `grep "bit$" /usr/share/dict/words`; do
  word='rabbit'
  wordbyte=`echo $word | sed -E 's/bit$/byte/'`

  image_url=`./main`

  wget $image_url -O temp.png
  convert temp.png -resize 200x temp.png
  magick montage -tile 1x temp.png top.png
  magick montage -tile 4x -title $wordbyte temp.png temp.png temp.png temp.png temp.png temp.png temp.png temp.png temp.png
  magick montage -tile 1x -title $word top.png temp.png -mode Concatenate output/$word.png
#done;