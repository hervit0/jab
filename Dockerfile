# FROM openjdk:7
# RUN mkdir -p opt/dynamodb
# WORKDIR /opt/dynamodb
# RUN wget https://s3.eu-central-1.amazonaws.com/dynamodb-local-frankfurt/dynamodb_local_latest.tar.gz -q -O - | tar -xz
# EXPOSE 8000
# ENTRYPOINT ["java", "-jar", "DynamoDBLocal.jar"]

FROM circleci/dynamodb
# RUN apt-get -y -qq update &&
#   apt-get -y -qq install python3.4-dev &&
#   curl -O https://bootstrap.pypa.io/get-pip.py &&
#   python3.4 get-pip.py --user &&
#   pip install awscli --upgrade --user
