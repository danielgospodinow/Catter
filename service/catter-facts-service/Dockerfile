FROM openjdk:11

WORKDIR /app

COPY . .

RUN ./gradlew clean build

CMD ["./gradlew", "bootRun"]
