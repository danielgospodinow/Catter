FROM openjdk:11
COPY . /usr/src/catter-facts-service
WORKDIR /usr/src/catter-facts-service
RUN ./gradlew clean build
CMD ["./gradlew", "bootRun"]