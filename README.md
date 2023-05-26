# stim-app-remover-cli
Stim App Remover CLI is a tool for removing built in stimulating and distracting apps from Android phones without the need to root the device. Apps such as the Play Store, Google Chrome Browser, YouTube, and YouTube Music are made uninstallable by usual means.  

Stim App remover can uninstall and reinstall these apps without the need for root access

## The apps
1. Play Store
2. Chrome Browser
3. YouTube
4. YouTube Music

## How To Use It
### 1. Setup The Device
1. Enable developer mode on the Android device
2. Enable USB Debugging on the device
3. Plug the device into the computer
4. Switch to File Transfer mode on the device

### 2. Build and Run
1. Run `make build`
2. Run `./stim-app-remover-cli`
3. A prompt may pop up on the device asking if you'd like to grant debug access to the plugged in computer. Grant access
3. Select either **Install** or **Uninstall** from the cli prompt

Note that the `adb` binary file must be in the same directory as the output binary file `stim-app-remover-cli`

## How It Works
Because the method utilized by Stim App Remover does not use root access, the apps are not completely removed from the pone. Instead they are _uninstalled_ for the main user of the phone but are still present on the system. They become unaccessible to the user and will not show up in the settings. The only way to re-instal them is through the adb cli tool or this application.

**Android Debug Bridge** or **adb** is the tool used to communicate with your phone. If you want to do things manually, you can run the adb commands yourself
**Install**  
`./adb shell pm uninstall --user 0 com.android.vending`  

**Uninstall**
`./adb shell cmd package install-existing com.android.vending`  

Just replace `com.android.vending` (Play Store apk name) with the _apk name_ of the app you want to add or remove

## Why This Was Made
Overstimulation from pointless web surfing, social media, endless scrolling on Instagram and TickTock, constantly having music playing etc... can result in a decreased ability to concentrate and remain focused for long periods of time, reduced cognitive ability, depression and anxiety, and wasted time.

How many people feel good about themselves after spending 3 hours scrolling through TickTock, binging on YouTube videos, watching porn, or surfing the web all day?  

By removing these distractions and putting up a barrier to reattaining them, individuals are forced to confront their thoughts and emotions, and are compelled to engage in activities that are more fulfilling and meaningful. Motivation and drive are also increased and one may suddenly find themselves driven to do things they previously had to force them selves to do.
