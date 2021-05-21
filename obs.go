package main

import (
	obsws "github.com/christopher-dG/go-obs-websocket"
)

type TextRequest struct {
	ItemName string `json:"itemName"`
	Text     string `json:"text"`
}

func updateOBSText(item, text string) error {
	getResp, err := obsws.NewGetTextGDIPlusPropertiesRequest(item).SendReceive(obsClient)
	if err != nil {
		return err
	}
	_, err = obsws.NewSetTextGDIPlusPropertiesRequest(
		getResp.Source,
		getResp.Align,
		getResp.BkColor,
		getResp.BkOpacity,
		getResp.Chatlog,
		getResp.ChatlogLines,
		getResp.Color,
		getResp.Extents,
		getResp.ExtentsCx,
		getResp.ExtentsCy,
		getResp.File,
		getResp.ReadFromFile,
		getResp.Font,
		getResp.FontFace,
		getResp.FontFlags,
		getResp.FontSize,
		getResp.FontStyle,
		getResp.Gradient,
		getResp.GradientColor,
		getResp.GradientDir,
		getResp.GradientOpacity,
		getResp.Outline,
		getResp.OutlineColor,
		getResp.OutlineSize,
		getResp.OutlineOpacity,
		text,
		getResp.Valign,
		getResp.Vertical,
		true,
	).SendReceive(obsClient)
	if err != nil {
		return err
	}
	return nil
}
