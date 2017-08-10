# galdric
machine learning server, for image classification


 - Reads all the datasets in the folder /dataset
 - Runs a server that allows to upload images to classify
 - Accepts PNG, JPG, JPEG
 - Each image is resized to the same size and converted to PNG type, configured in the config.json
 - For the input images, calculates the euclidean distances
 - Applyies KNN (K-Nearest Neighbours algorithm) to classify the images
 - Server returns the classification result, that is the label of the object in the image



### Instructions

Put dataset in /dataset directory with subdirectories, where each subdirectory contains images of one element.

For example:
```
dataset/
    leopard/
        img01.png
        img02.png
        img03.png
        ...
    laptop/
        img01.png
        img02.png
        ...
    camera/
        img01.png
        img02.png
        ...
```
So, we have each image and to which element category is (the name of subdirectory).


Then, run the server:
```
    >./galdric
```

Now, just need to perform petitions with new images, to get the response from the server classifying them:
```bash
    curl -F file=@./testimage.png http://127.0.0.1:3055/image
```
And the server will return:
```
    seems to be a leopard
```

Can perform some tests with the test.sh file:
```
    bash test.sh
```


-------------
#### Useful commands

send file over ssh:
```
scp dataset.tar.gz root@SERVERIP:/root/galdric
```

on the server, untar file:
```
tar -xvzf dataset.tar.gz
```
