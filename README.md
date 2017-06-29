# galdric
machine learning server, for image classification



 - Reads all the datasets in the folder /dataset
 - Each image is resized to the same size, configured in the config.json
 - For the input images, calculates the euclidean distances
 - Gets the nearest neighbour
 - Show the result, that is the label of the object in the image













-------------

send file over ssh:
```
scp dataset.tar.gz root@51.255.193.106:/root/galdric
```

on the server, untar file:
```
tar -xvzf dataset.tar.gz
```