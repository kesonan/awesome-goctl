// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.User.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.response.HttpResponseData;

public class LoginResp extends HttpResponseData {

	private String token = "";
	private String uid = "";

	public LoginResp() {
	}

	public LoginResp(String token, String uid) {
		this.token = token;
		this.uid = uid;
	}

	@NotNull 
	public String getToken() {
		return this.token;
	}

	public void setToken(@NotNull String token) {
		this.token = token;
	}

	@NotNull 
	public String getUid() {
		return this.uid;
	}

	public void setUid(@NotNull String uid) {
		this.uid = uid;
	}

}