package launcher

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/sqweek/dialog"
	"github.com/wieku/danser-go/app/utils"
	"github.com/wieku/danser-go/framework/platform"
	"strings"
)

type messageType int

const (
	mInfo = messageType(iota)
	mError
	mQuestion
)

func showMessage(typ messageType, format string, args ...any) bool {
	message := fmt.Sprintf(format, args...)

	switch typ {
	case mInfo:
		dialog.Message(message).Info()
	case mError:
		if urlIndex := strings.Index(message, "http"); urlIndex > -1 {
			if dialog.Message(message + "\n\nDo you want to go there?").ErrorYesNo() {
				url := message[urlIndex:]
				platform.OpenURL(url)
				return true
			}
		} else {
			dialog.Message(message).Error()
		}
	case mQuestion:
		return dialog.Message(message).YesNo()
	}

	return false
}

func checkForUpdates(pingUpToDate bool) {
	status, url, err := utils.CheckForUpdate()

	switch status {
	case utils.Ignored, utils.UpToDate:
		if pingUpToDate {
			showMessage(mInfo, "You're using the newest version of danser.")
		}
	case utils.Failed:
		showMessage(mError, "Can't get version from GitHub:", err)
	case utils.Snapshot:
		if showMessage(mQuestion, "You're using a snapshot version of danser.\nFor newer version of snapshots please visit the official danser discord server at: %s\n\nDo you want to go there?", url) {
			platform.OpenURL(url)
		}
	case utils.UpdateAvailable:
		if showMessage(mQuestion, "You're using an older version of danser.\nYou can download a newer version here: %s\n\nDo you want to go there?", url) {
			platform.OpenURL(url)
		}
	}
}

func vec2(x, y float32) imgui.Vec2 {
	return imgui.Vec2{
		X: x,
		Y: y,
	}
}

func vec4(x, y, z, w float32) imgui.Vec4 {
	return imgui.Vec4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func vzero() imgui.Vec2 {
	return vec2(0, 0)
}
