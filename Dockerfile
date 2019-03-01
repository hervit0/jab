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
ARG CLI_VERSION=1.16.86

ENV AWS_ACCESS_KEY_ID="replace"
ENV AWS_SECRET_ACCESS_KEY="replace"
ENV AWS_DEFAULT_REGION="us-west-2"
ENV AWS_DEFAULT_OUTPUT="json"

RUN apt-get update && apt-get install -y python-pip

RUN pip install --no-cache-dir awscli==$CLI_VERSION

RUN aws dynamodb create-table --endpoint-url http://127.0.0.1:8000 --table-name fraud-whitelist --key-schema AttributeName=Cli,KeyType=HASH AttributeName=TenancyId,KeyType=RANGE --attribute-definitions AttributeName=Cli,AttributeType=S AttributeName=TenancyId,AttributeType=S --billing-mode PAY_PER_REQUEST
