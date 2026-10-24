package salamander

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"
)

type zeroSource struct{}

func (zeroSource) Int63() int64 { return 0 }
func (zeroSource) Seed(int64)   {}
func TestOriginalObfuscationMachineGoldens(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 8)
		n := o.Obfuscate(input, output)
		if n != 8 || hex.EncodeToString(output) != "0000000000000000" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("1", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("2", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("03")
		output := bytes.Repeat([]byte{0xa5}, 9)
		n := o.Obfuscate(input, output)
		if n != 9 || hex.EncodeToString(output) != "0000000000000000fa" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("3", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa")
		output := bytes.Repeat([]byte{0xa5}, 1)
		n := o.Deobfuscate(input, output)
		if n != 1 || hex.EncodeToString(output) != "03" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("4", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5")
		output := bytes.Repeat([]byte{0xa5}, 39)
		n := o.Obfuscate(input, output)
		if n != 39 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("5", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a")
		output := bytes.Repeat([]byte{0xa5}, 31)
		n := o.Deobfuscate(input, output)
		if n != 31 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("6", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc")
		output := bytes.Repeat([]byte{0xa5}, 40)
		n := o.Obfuscate(input, output)
		if n != 40 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("7", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23")
		output := bytes.Repeat([]byte{0xa5}, 32)
		n := o.Deobfuscate(input, output)
		if n != 32 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("8", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3")
		output := bytes.Repeat([]byte{0xa5}, 41)
		n := o.Obfuscate(input, output)
		if n != 41 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("9", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a")
		output := bytes.Repeat([]byte{0xa5}, 33)
		n := o.Deobfuscate(input, output)
		if n != 33 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("10", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4")
		output := bytes.Repeat([]byte{0xa5}, 2048)
		n := o.Obfuscate(input, output)
		if n != 2048 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("11", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac")
		output := bytes.Repeat([]byte{0xa5}, 2040)
		n := o.Deobfuscate(input, output)
		if n != 2040 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("12", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("13", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("00000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("14", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("15", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("000000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("16", func(t *testing.T) {
		key := make([]byte, 23, 23+0)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("31323334")
		output := bytes.Repeat([]byte{0xa5}, 11)
		n := o.Obfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "a5a5a5a5a5a5a5a5a5a5a5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b6579" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("17", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 8)
		n := o.Obfuscate(input, output)
		if n != 8 || hex.EncodeToString(output) != "0000000000000000" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("18", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("19", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("03")
		output := bytes.Repeat([]byte{0xa5}, 9)
		n := o.Obfuscate(input, output)
		if n != 9 || hex.EncodeToString(output) != "0000000000000000fa" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("20", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa")
		output := bytes.Repeat([]byte{0xa5}, 1)
		n := o.Deobfuscate(input, output)
		if n != 1 || hex.EncodeToString(output) != "03" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("21", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5")
		output := bytes.Repeat([]byte{0xa5}, 39)
		n := o.Obfuscate(input, output)
		if n != 39 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("22", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a")
		output := bytes.Repeat([]byte{0xa5}, 31)
		n := o.Deobfuscate(input, output)
		if n != 31 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("23", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc")
		output := bytes.Repeat([]byte{0xa5}, 40)
		n := o.Obfuscate(input, output)
		if n != 40 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("24", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23")
		output := bytes.Repeat([]byte{0xa5}, 32)
		n := o.Deobfuscate(input, output)
		if n != 32 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("25", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3")
		output := bytes.Repeat([]byte{0xa5}, 41)
		n := o.Obfuscate(input, output)
		if n != 41 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("26", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a")
		output := bytes.Repeat([]byte{0xa5}, 33)
		n := o.Deobfuscate(input, output)
		if n != 33 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("27", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4")
		output := bytes.Repeat([]byte{0xa5}, 2048)
		n := o.Obfuscate(input, output)
		if n != 2048 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("28", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac")
		output := bytes.Repeat([]byte{0xa5}, 2040)
		n := o.Deobfuscate(input, output)
		if n != 2040 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65790000000000000000" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("29", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("30", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("00000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("31", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("32", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("000000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("33", func(t *testing.T) {
		key := make([]byte, 23, 23+8)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("31323334")
		output := bytes.Repeat([]byte{0xa5}, 11)
		n := o.Obfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "a5a5a5a5a5a5a5a5a5a5a5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("34", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 8)
		n := o.Obfuscate(input, output)
		if n != 8 || hex.EncodeToString(output) != "0000000000000000" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("35", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("36", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("03")
		output := bytes.Repeat([]byte{0xa5}, 9)
		n := o.Obfuscate(input, output)
		if n != 9 || hex.EncodeToString(output) != "0000000000000000fa" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("37", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa")
		output := bytes.Repeat([]byte{0xa5}, 1)
		n := o.Deobfuscate(input, output)
		if n != 1 || hex.EncodeToString(output) != "03" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("38", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5")
		output := bytes.Repeat([]byte{0xa5}, 39)
		n := o.Obfuscate(input, output)
		if n != 39 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("39", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a")
		output := bytes.Repeat([]byte{0xa5}, 31)
		n := o.Deobfuscate(input, output)
		if n != 31 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("40", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc")
		output := bytes.Repeat([]byte{0xa5}, 40)
		n := o.Obfuscate(input, output)
		if n != 40 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("41", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a23")
		output := bytes.Repeat([]byte{0xa5}, 32)
		n := o.Deobfuscate(input, output)
		if n != 32 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dc" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("42", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3")
		output := bytes.Repeat([]byte{0xa5}, 41)
		n := o.Obfuscate(input, output)
		if n != 41 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("43", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a")
		output := bytes.Repeat([]byte{0xa5}, 33)
		n := o.Deobfuscate(input, output)
		if n != 33 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("44", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4")
		output := bytes.Repeat([]byte{0xa5}, 2048)
		n := o.Obfuscate(input, output)
		if n != 2048 || hex.EncodeToString(output) != "0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("45", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9acb6950e5fea10aa03fa74ad6c17b6fda2a6ddce68b87ccbb15f4b92058137e9ccd6f56e7fca308a231a944d8cf796dd8286bdae08d81ceb917f6b72e561d709ecf6d54e1faa50ea433ab46dacd7763d62669d8e28f83c8bf11f0b52c541f7290c1635ae3f8a70ca635ad40dccb7561d42467d6ec818dcabd13f2b32a52197492c36158edf6a902a837af42dec97367d22265d4ee838fc4b31dfcb128501b7694c5675eeff4ab00aa39a14cd0c77165d02063d2e88589c6b11ffebf265e157896c7655ce9f2ad06ac3ba34ed2c57f6bde2e61d0ea878bc0b719f8bd245c177a98c96b52ebf0af04ae3da548d4c37d69dc2c6fdee48985c2b51bfabb225a117c9ac")
		output := bytes.Repeat([]byte{0xa5}, 2040)
		n := o.Deobfuscate(input, output)
		if n != 2040 || hex.EncodeToString(output) != "030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4cbd2d9e0e7eef5fc030a11181f262d343b424950575e656c737a81888f969da4abb2b9c0c7ced5dce3eaf1f8ff060d141b222930373e454c535a61686f767d848b9299a0a7aeb5bcc3cad1d8dfe6edf4fb020910171e252c333a41484f565d646b727980878e959ca3aab1b8bfc6cdd4dbe2e9f0f7fe050c131a21282f363d444b525960676e757c838a91989fa6adb4bbc2c9d0d7dee5ecf3fa01080f161d242b323940474e555c636a71787f868d949ba2a9b0b7bec5ccd3dae1e8eff6fd040b121920272e353c434a51585f666d747b828990979ea5acb3bac1c8cfd6dde4ebf2f900070e151c232a31383f464d545b626970777e858c939aa1a8afb6bdc4" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b657900000000000000005a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("46", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("47", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("00000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("48", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("0000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("49", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("000000000000000000")
		output := bytes.Repeat([]byte{0xa5}, 0)
		n := o.Deobfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
	t.Run("50", func(t *testing.T) {
		key := make([]byte, 23, 23+32)
		copy(key, "original-salamander-key")
		backing := key[:cap(key)]
		for i := 23; i < len(backing); i++ {
			backing[i] = 0x5a
		}
		o := DOqsrH8iC{EDkKnfGR8w: key, FY8Ki3jl: rand.New(zeroSource{})}
		input, _ := hex.DecodeString("31323334")
		output := bytes.Repeat([]byte{0xa5}, 11)
		n := o.Obfuscate(input, output)
		if n != 0 || hex.EncodeToString(output) != "a5a5a5a5a5a5a5a5a5a5a5" || hex.EncodeToString(backing) != "6f726967696e616c2d73616c616d616e6465722d6b65795a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a5a" {
			t.Fatalf("n=%d output=%x backing=%x", n, output, backing)
		}
	})
}
