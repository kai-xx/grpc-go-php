<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace GPBMetadata;

class User
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?

user.protouser"\'

UserEntity
name (	
age ("3
UserIndexRequest
page (
	page_size ("M
UserIndexResponse
err (
msg (	
data (2.user.UserEntity"
UserViewRequest
uid ("L
UserViewResponse
err (
msg (	
data (2.user.UserEntity">
UserPostRequest
name (	
password (	
age (",
UserPostResponse
err (
msg (	" 
UserDeleteRequest
uid (".
UserDeleteResponse
err (
msg (	*9
EnumUserSex
SEX_INIT 
SEX_MALE

SEX_FEMALE2?
User>
	UserIndex.user.UserIndexRequest.user.UserIndexResponse" ;
UserView.user.UserViewRequest.user.UserViewResponse" ;
UserPost.user.UserPostRequest.user.UserPostResponse" A

UserDelete.user.UserDeleteRequest.user.UserDeleteResponse" bproto3'
        , true);

        static::$is_initialized = true;
    }
}

