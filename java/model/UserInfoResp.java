// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.User.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.response.HttpResponseData;
import com.xhb.core.packet.HttpData;

public class UserInfoResp extends HttpResponseData {

	private String name = "";
	private String email = "";

	public UserInfoResp() {
	}

	public UserInfoResp(String name, String email) {
		this.name = name;
		this.email = email;
	}

	@NotNull 
	public String getName() {
		return this.name;
	}

	public void setName(@NotNull String name) {
		this.name = name;
	}

	@NotNull 
	public String getEmail() {
		return this.email;
	}

	public void setEmail(@NotNull String email) {
		this.email = email;
	}

}
