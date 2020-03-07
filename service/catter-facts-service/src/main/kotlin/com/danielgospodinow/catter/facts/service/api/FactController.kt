package com.danielgospodinow.catter.facts.service.api

import com.danielgospodinow.catter.facts.service.model.Fact
import com.danielgospodinow.catter.facts.service.service.FactService
import org.springframework.web.bind.annotation.*
import reactor.core.publisher.Mono

@RestController
@RequestMapping("/api/fact")
class FactController(
        val factService: FactService) {

    @GetMapping("/{factId}")
    fun getFact(@PathVariable factId: String): Mono<Fact> {
        return Mono.empty()
    }

    @GetMapping("/random")
    fun getRandomFact(): Mono<Fact> {
        return Mono.empty()
    }

    @PostMapping
    fun createFact(factMono: Mono<Fact>): Mono<Fact> {
        return Mono.empty()
    }

    @DeleteMapping("/{factId}")
    fun deleteFact(@PathVariable factId: String) {

    }
}