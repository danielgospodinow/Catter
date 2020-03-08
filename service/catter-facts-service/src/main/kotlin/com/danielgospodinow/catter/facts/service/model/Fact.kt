package com.danielgospodinow.catter.facts.service.model

import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document

@Document(value = "facts")
data class Fact(
        @Id var id: String?,
        var content: String)