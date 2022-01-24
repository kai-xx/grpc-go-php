<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace User;

use UnexpectedValueException;

/**
 * 枚举类型
 *
 * Protobuf type <code>user.EnumUserSex</code>
 */
class EnumUserSex
{
    /**
     * 枚举类型必须以 0 起始
     *
     * Generated from protobuf enum <code>SEX_INIT = 0;</code>
     */
    const SEX_INIT = 0;
    /**
     * Generated from protobuf enum <code>SEX_MALE = 1;</code>
     */
    const SEX_MALE = 1;
    /**
     * Generated from protobuf enum <code>SEX_FEMALE = 2;</code>
     */
    const SEX_FEMALE = 2;

    private static $valueToName = [
        self::SEX_INIT => 'SEX_INIT',
        self::SEX_MALE => 'SEX_MALE',
        self::SEX_FEMALE => 'SEX_FEMALE',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

