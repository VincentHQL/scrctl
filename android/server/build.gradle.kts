plugins {
    alias(libs.plugins.android.application)
}

android {
    namespace = "com.vincenthql.scrctl"
    compileSdk = 34
    defaultConfig {
        applicationId = "com.vincenthql.scrctl"
        minSdk = 24
        targetSdk = 34
        versionCode = 1
        versionName = "1.0"
        testInstrumentationRunner = "android.support.test.runner.AndroidJUnitRunner"
    }
    buildTypes {
        release {
            isMinifyEnabled = false
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro"
            )
        }
    }
    buildFeatures {
        buildConfig = true
        aidl = true
    }
    applicationVariants.all {
        outputs.all {
        }
    }
}

dependencies {
    testImplementation(libs.junit)
}

