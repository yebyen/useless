package io.mecris.vindbox.api

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json
import io.mecris.vindbox.BuildConfig

class VindBoxClient(
    private val baseUrl: String = BuildConfig.API_URL
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
