image=chenchun/stress && docker build --network=host -t $image . && docker push $image
