package com.danielgospodinow.catter.facts.service.service

import com.danielgospodinow.catter.facts.service.repository.FactRepository
import org.springframework.stereotype.Service

@Service
class FactService(
        val factRepository: FactRepository) {


}