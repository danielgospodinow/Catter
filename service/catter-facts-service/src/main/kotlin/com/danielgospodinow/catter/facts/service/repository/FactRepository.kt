package com.danielgospodinow.catter.facts.service.repository

import com.danielgospodinow.catter.facts.service.model.Fact
import org.springframework.data.mongodb.repository.ReactiveMongoRepository
import org.springframework.stereotype.Repository

@Repository
interface FactRepository : ReactiveMongoRepository<Fact, String> {


}