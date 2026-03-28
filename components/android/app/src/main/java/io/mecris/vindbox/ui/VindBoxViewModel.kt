package io.mecris.vindbox.ui

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import io.mecris.vindbox.api.UselessMachineStatus
import io.mecris.vindbox.api.VindBoxClient
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class VindBoxViewModel : ViewModel() {
    private val client = VindBoxClient()

    private val _status = MutableStateFlow<UselessMachineStatus?>(null)
    val status: StateFlow<UselessMachineStatus?> = _status.asStateFlow()

    private val _error = MutableStateFlow<String?>(null)
    val error: StateFlow<String?> = _error.asStateFlow()

    init {
        startPolling()
    }

    private fun startPolling() {
        viewModelScope.launch {
            while (true) {
                refreshStatus()
                delay(5000) // Poll every 5 seconds
            }
        }
    }

    fun refreshStatus() {
        viewModelScope.launch {
            client.getStatus()
                .onSuccess { 
                    _status.value = it
                    _error.value = null
                }
                .onFailure { _error.value = "Failed to fetch status: ${it.message}" }
        }
    }

    fun pushButton() {
        viewModelScope.launch {
            client.pushButton()
                .onSuccess { 
                    _error.value = null
                    refreshStatus()
                }
                .onFailure { _error.value = "Failed to push button: ${it.message}" }
        }
    }

    override fun onCleared() {
        super.onCleared()
        client.close()
    }
}
