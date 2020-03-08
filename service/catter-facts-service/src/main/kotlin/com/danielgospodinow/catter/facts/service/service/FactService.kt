package com.danielgospodinow.catter.facts.service.service

import com.danielgospodinow.catter.facts.service.model.Fact
import com.danielgospodinow.catter.facts.service.repository.FactRepository
import org.springframework.stereotype.Service
import reactor.core.publisher.Mono

@Service
class FactService(
        private val factRepository: FactRepository) {

    fun getFact(factId: String): Mono<Fact> = factRepository.findById(factId)

    fun getRandomFact(): Mono<Fact> = factRepository.findAll().last()

    fun createFact(fact: Fact): Mono<Fact> = factRepository.save(fact)

    fun deleteFact(factId: String) = factRepository.deleteById(factId)
}