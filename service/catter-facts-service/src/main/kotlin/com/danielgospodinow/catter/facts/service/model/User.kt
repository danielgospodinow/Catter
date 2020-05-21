package com.danielgospodinow.catter.facts.service.model

import org.springframework.data.mongodb.core.mapping.Document

@Document("users")
data class User(
        var id: String)
