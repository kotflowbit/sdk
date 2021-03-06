<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rpc/block.proto

namespace KotFlowBlock;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>KotFlowBlock.Header</code>
 */
class Header extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bytes ParentHash = 1;</code>
     */
    protected $ParentHash = '';
    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     */
    protected $Hash = '';
    /**
     * Generated from protobuf field <code>bytes Time = 3;</code>
     */
    protected $Time = '';
    /**
     * Generated from protobuf field <code>bytes Nonce = 4;</code>
     */
    protected $Nonce = '';
    /**
     * Generated from protobuf field <code>bytes Miner = 5;</code>
     */
    protected $Miner = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $ParentHash
     *     @type string $Hash
     *     @type string $Time
     *     @type string $Nonce
     *     @type string $Miner
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Rpc\Block::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bytes ParentHash = 1;</code>
     * @return string
     */
    public function getParentHash()
    {
        return $this->ParentHash;
    }

    /**
     * Generated from protobuf field <code>bytes ParentHash = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setParentHash($var)
    {
        GPBUtil::checkString($var, False);
        $this->ParentHash = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     * @return string
     */
    public function getHash()
    {
        return $this->Hash;
    }

    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setHash($var)
    {
        GPBUtil::checkString($var, False);
        $this->Hash = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes Time = 3;</code>
     * @return string
     */
    public function getTime()
    {
        return $this->Time;
    }

    /**
     * Generated from protobuf field <code>bytes Time = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setTime($var)
    {
        GPBUtil::checkString($var, False);
        $this->Time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes Nonce = 4;</code>
     * @return string
     */
    public function getNonce()
    {
        return $this->Nonce;
    }

    /**
     * Generated from protobuf field <code>bytes Nonce = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setNonce($var)
    {
        GPBUtil::checkString($var, False);
        $this->Nonce = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes Miner = 5;</code>
     * @return string
     */
    public function getMiner()
    {
        return $this->Miner;
    }

    /**
     * Generated from protobuf field <code>bytes Miner = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setMiner($var)
    {
        GPBUtil::checkString($var, False);
        $this->Miner = $var;

        return $this;
    }

}

