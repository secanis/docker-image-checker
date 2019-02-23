# reallife example to check a base image in your CI/CD pipeline for i.e.
if docker run --rm secanis/docker-image-checker:latest \
        ./dic -base "hub.docker.com/library/alpine:latest" \
        -image "hub.docker.com/secanis/stjorna:2.1.3"; then
    # do something in your pipeline
    echo "Command succeeded"
else
    # do something in your pipeline
    echo "Command failed"
fi