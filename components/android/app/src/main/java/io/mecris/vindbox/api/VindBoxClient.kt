package io.mecris.vindbox.api

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json

class VindBoxClient(
    private val baseUrl: String = "http://10.0.2.2:3000"
) {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json {
                ignoreUnknownKeys = true
                coerceInputValues = true
            })
        }
    }

    suspend fun getStatus(): Result<UselessMachineStatus> = runCatching {
        client.get("$baseUrl/status").body()
    }

    suspend fun pushButton(): Result<String> = runCatching {
        client.post("$baseUrl/push").body()
    }

    fun close() {
        client.close()
    }
}
