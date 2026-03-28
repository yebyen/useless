package io.mecris.vindbox

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.activity.viewModels
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import io.mecris.vindbox.ui.VindBoxViewModel
import io.mecris.vindbox.ui.theme.VindBoxTheme

class MainActivity : ComponentActivity() {
    private val viewModel: VindBoxViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            VindBoxTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    UselessMachineScreen(
                        viewModel = viewModel,
                        modifier = Modifier.padding(innerPadding)
                    )
                }
            }
        }
    }
}

@Composable
fun UselessMachineScreen(
    viewModel: VindBoxViewModel,
    modifier: Modifier = Modifier
) {
    val status by viewModel.status.collectAsState()
    val error by viewModel.error.collectAsState()

    val backgroundColor = if (status?.isNagging == true) {
        Color(0xFFFFEBEE) // Light red when nagging
    } else {
        MaterialTheme.colorScheme.background
    }

    Column(
        modifier = modifier
            .fillMaxSize()
            .background(backgroundColor)
            .padding(16.dp),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.Center
    ) {
        Text(
            text = "Vind-Box",
            fontSize = 32.sp,
            fontWeight = FontWeight.Bold,
            color = MaterialTheme.colorScheme.primary
        )
        
        Spacer(modifier = Modifier.height(48.dp))

        if (status != null) {
            Text(
                text = "Daily Pushes: ${status!!.dailyCount}",
                fontSize = 24.sp
            )
            Text(
                text = "Last Pushed: ${status!!.lastPushed}",
                fontSize = 14.sp,
                color = Color.Gray
            )
            
            if (status!!.isNagging) {
                Spacer(modifier = Modifier.height(16.dp))
                Text(
                    text = "Don't be lazy! Push it!",
                    color = Color.Red,
                    fontWeight = FontWeight.Bold
                )
            }
        } else {
            CircularProgressIndicator()
            Text(text = "Connecting to Brain...")
        }

        Spacer(modifier = Modifier.height(64.dp))

        Button(
            onClick = { viewModel.pushButton() },
            modifier = Modifier
                .size(200.dp),
            shape = MaterialTheme.shapes.extraLarge
        ) {
            Text(
                text = "PUSH",
                fontSize = 24.sp,
                fontWeight = FontWeight.Black
            )
        }

        if (error != null) {
            Spacer(modifier = Modifier.height(32.dp))
            Text(
                text = error!!,
                color = Color.Red,
                fontSize = 12.sp
            )
            Button(onClick = { viewModel.refreshStatus() }) {
                Text("Retry")
            }
        }
    }
}
