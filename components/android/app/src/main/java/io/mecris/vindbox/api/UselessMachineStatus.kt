package io.mecris.vindbox.api

import kotlinx.serialization.Serializable

@Serializable
data class UselessMachineStatus(
    val dailyCount: Int,
    val lastPushed: String,
    val isNagging: Boolean
)
