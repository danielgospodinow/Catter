package com.danielgospodinow.catter.facts.service

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class CatterFactsServiceApplication

fun main(args: Array<String>) {
	runApplication<CatterFactsServiceApplication>(*args)
}
