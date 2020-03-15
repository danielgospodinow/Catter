package com.danielgospodinow.catter.facts.service.api

import com.danielgospodinow.catter.facts.service.model.Fact
import com.danielgospodinow.catter.facts.service.service.FactService
import org.springframework.web.bind.annotation.*
import reactor.core.publisher.Mono

@RestController
@RequestMapping("/api/fact")
class FactController(
        private val factService: FactService) {

    @GetMapping("/{factId}")
    fun getFact(@PathVariable factId: String): Mono<Fact> = factService.getFact(factId)

    @GetMapping("/random")
    fun getRandomFact(): Mono<Fact> = factService.getRandomFact()

    @PostMapping
    fun createFact(@RequestBody factMono: Mono<Fact>): Mono<Fact> = factMono.flatMap(factService::createFact)

    @DeleteMapping("/{factId}")
    fun deleteFact(@PathVariable factId: String): Mono<Void> = factService.deleteFact(factId)
}