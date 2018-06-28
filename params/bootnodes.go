// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",

	// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	// "enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	// "enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DccsBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Dccs network.
var DccsBootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",
}
