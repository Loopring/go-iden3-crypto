package babyjub

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"

	"github.com/iden3/go-iden3-crypto/constants"
	"github.com/iden3/go-iden3-crypto/ff"
	"github.com/iden3/go-iden3-crypto/utils"
	"github.com/stretchr/testify/assert"
)

func zero() *ff.Element {
	return ff.NewElement().SetZero()
}
func one() *ff.Element {
	return ff.NewElement().SetOne()
}

func TestAdd1(t *testing.T) {
	a := &Point{X: zero(), Y: one()}
	b := &Point{X: zero(), Y: one()}

	c := NewPoint().Add(a, b)
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t, "0", c.X.String())
	assert.Equal(t, "1", c.Y.String())
}

func TestAdd2(t *testing.T) {
	aX := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	aY := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	a := &Point{X: aX, Y: aY}

	bX := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	bY := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Add(a, b)
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"6890855772600357754907169075114257697580319025794532037257385534741338397365",
		c.X.String())
	assert.Equal(t,
		"4338620300185947561074059802482547481416142213883829469920100239455078257889",
		c.Y.String())
}

func TestAdd3(t *testing.T) {
	aX := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	aY := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	a := &Point{X: aX, Y: aY}

	bX := ff.NewElement().SetString(
		"16540640123574156134436876038791482806971768689494387082833631921987005038935")
	bY := ff.NewElement().SetString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Add(a, b)
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"7916061937171219682591368294088513039687205273691143098332585753343424131937",
		c.X.String())
	assert.Equal(t,
		"14035240266687799601661095864649209771790948434046947201833777492504781204499",
		c.Y.String())
}

func TestAdd4(t *testing.T) {
	aX := ff.NewElement().SetString(
		"0")
	aY := ff.NewElement().SetString(
		"1")
	a := &Point{X: aX, Y: aY}

	bX := ff.NewElement().SetString(
		"16540640123574156134436876038791482806971768689494387082833631921987005038935")
	bY := ff.NewElement().SetString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	b := &Point{X: bX, Y: bY}

	c := NewPoint().Add(a, b)
	// fmt.Printf("%v = 2 * %v", *c, *a)
	assert.Equal(t,
		"16540640123574156134436876038791482806971768689494387082833631921987005038935",
		c.X.String())
	assert.Equal(t,
		"20819045374670962167435360035096875258406992893633759881276124905556507972311",
		c.Y.String())
}

func TestInCurve1(t *testing.T) {
	p := &Point{X: zero(), Y: one()}
	assert.Equal(t, true, p.InCurve())
}

func TestInCurve2(t *testing.T) {
	p := &Point{X: one(), Y: zero()}
	assert.Equal(t, false, p.InCurve())
}

func TestMul0(t *testing.T) {
	x := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString("3")

	r2 := NewPoint().Add(p, p)
	r2 = NewPoint().Add(r2, p)
	r := NewPoint().Mul(s, p)
	assert.Equal(t, r2.X.String(), r.X.String())
	assert.Equal(t, r2.Y.String(), r.Y.String())

	assert.Equal(t,
		"19372461775513343691590086534037741906533799473648040012278229434133483800898",
		r.X.String())
	assert.Equal(t,
		"9458658722007214007257525444427903161243386465067105737478306991484593958249",
		r.Y.String())
}

func TestMul1(t *testing.T) {
	x := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString(
		"14035240266687799601661095864649209771790948434046947201833777492504781204499")
	r := NewPoint().Mul(s, p)
	assert.Equal(t,
		"17070357974431721403481313912716834497662307308519659060910483826664480189605",
		r.X.String())
	assert.Equal(t,
		"4014745322800118607127020275658861516666525056516280575712425373174125159339",
		r.Y.String())
}

func TestMul2(t *testing.T) {
	x := ff.NewElement().SetString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := ff.NewElement().SetString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	s := utils.NewIntFromString(
		"20819045374670962167435360035096875258406992893633759881276124905556507972311")
	r := NewPoint().Mul(s, p)
	assert.Equal(t,
		"13563888653650925984868671744672725781658357821216877865297235725727006259983",
		r.X.String())
	assert.Equal(t,
		"8442587202676550862664528699803615547505326611544120184665036919364004251662",
		r.Y.String())
}

func TestInCurve3(t *testing.T) {
	x := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InCurve())
}

func TestInCurve4(t *testing.T) {
	x := ff.NewElement().SetString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := ff.NewElement().SetString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InCurve())
}

func TestInSubGroup1(t *testing.T) {
	x := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InSubGroup())
}

func TestInSubGroup2(t *testing.T) {
	x := ff.NewElement().SetString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := ff.NewElement().SetString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}
	assert.Equal(t, true, p.InSubGroup())
}

func TestCompressDecompress1(t *testing.T) {
	x := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	y := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	p := &Point{X: x, Y: y}

	buf := p.Compress()
	assert.Equal(t, "53b81ed5bffe9545b54016234682e7b2f699bd42a5e9eae27ff4051bc698ce85", hex.EncodeToString(buf[:]))

	p2, err := NewPoint().Decompress(buf)
	assert.Equal(t, nil, err)
	assert.Equal(t, p.X.String(), p2.X.String())
	assert.Equal(t, p.Y.String(), p2.Y.String())
}

func TestCompressDecompress2(t *testing.T) {
	x := ff.NewElement().SetString(
		"6890855772600357754907169075114257697580319025794532037257385534741338397365")
	y := ff.NewElement().SetString(
		"4338620300185947561074059802482547481416142213883829469920100239455078257889")
	p := &Point{X: x, Y: y}

	buf := p.Compress()
	assert.Equal(t, "e114eb17eddf794f063a68fecac515e3620e131976108555735c8b0773929709", hex.EncodeToString(buf[:]))

	p2, err := NewPoint().Decompress(buf)
	assert.Equal(t, nil, err)
	assert.Equal(t, p.X.String(), p2.X.String())
	assert.Equal(t, p.Y.String(), p2.Y.String())
}

func TestCompressDecompressRnd(t *testing.T) {
	for i := 0; i < 64; i++ {
		p1 := NewPoint().Mul(big.NewInt(int64(i)), B8)
		buf := p1.Compress()
		p2, err := NewPoint().Decompress(buf)
		assert.Equal(t, nil, err)
		// assert.Equal(t, p1, p2)
		assert.True(t, p1.Equal(p2))
	}
}

func BenchmarkBabyjub(b *testing.B) {
	const n = 256

	rnd := rand.New(rand.NewSource(42))

	var badpoints [n]*Point
	for i := 0; i < n; i++ {
		x := ff.NewElement().SetRandom()
		y := ff.NewElement().SetRandom()
		badpoints[i] = &Point{X: x, Y: y}
	}

	var points [n]*Point
	baseX := ff.NewElement().SetString(
		"17777552123799933955779906779655732241715742912184938656739573121738514868268")
	baseY := ff.NewElement().SetString(
		"2626589144620713026669568689430873010625803728049924121243784502389097019475")
	base := &Point{X: baseX, Y: baseY}
	for i := 0; i < n; i++ {
		s := new(big.Int).Rand(rnd, constants.Q)
		points[i] = NewPoint().Mul(s, base)
	}

	var scalars [n]*big.Int
	for i := 0; i < n; i++ {
		scalars[i] = new(big.Int).Rand(rnd, constants.Q)
	}

	b.Run("AddConst", func(b *testing.B) {
		p0 := &Point{X: zero(), Y: one()}
		p1 := &Point{X: zero(), Y: one()}

		p2 := NewPoint()
		for i := 0; i < b.N; i++ {
			p2.Add(p0, p1)
		}
	})

	b.Run("AddRnd", func(b *testing.B) {
		res := NewPoint()
		for i := 0; i < b.N; i++ {
			res.Add(points[i%(n/2)], points[i%(n/2)+1])
		}
	})

	b.Run("MulRnd", func(b *testing.B) {
		res := NewPoint()
		for i := 0; i < b.N; i++ {
			res.Mul(scalars[i%n], points[i%n])
		}
	})

	b.Run("Compress", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[i%n].Compress()
		}
	})

	b.Run("InCurve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			badpoints[i%n].InCurve()
		}
	})

	b.Run("InSubGroup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			points[i%n].InCurve()
		}
	})
}
