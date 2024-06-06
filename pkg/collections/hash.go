/*
 *  Copyright 2024 F5, Inc. All rights reserved.
 *
 *  No part of the software may be reproduced or transmitted in any
 *  form or by any means, electronic or mechanical, for any purpose,
 *  without express written permission of F5 Networks, Inc.
 *
 */

package collections

import (
	"crypto/sha512"
	"encoding/hex"
)

// Hash the provided data using sha512. This is a consistent hash.
func Hash(data []byte) string {
	b := sha512.Sum512_224(data)
	return hex.EncodeToString(b[:])
}
