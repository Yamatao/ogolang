package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("no letters - no words", func(t *testing.T) {
		require.Len(t, Top10(` 324 #$5 @#^, .><:} }{ \" ; \' *( /)@!#0$ # $ % ^ *( !-~=`), 0)
	})

	t.Run("one word", func(t *testing.T) {
		result := Top10("abc")
		require.Equal(t, result, []string{"abc"})
	})

	t.Run("punctuation does not matter", func(t *testing.T) {
		result := Top10(" a,    a    .b! c c . ")
		expected := []string{"a", "b", "c"}
		require.ElementsMatch(t, expected, result)
	})

	t.Run("case insensitivity", func(t *testing.T) {
		result := Top10("Aaa aaA bBb BBB")
		require.ElementsMatch(t, []string{"aaa", "bbb"}, result)
	})

	t.Run("frequency", func(t *testing.T) {
		result := Top10("a b c d e f g h i j X X Y Y k l m n o p q r s t v u w")
		require.Subset(t, result, []string{"x", "y"})
		require.Len(t, result, 10)
	})

	t.Run("sorted order", func(t *testing.T) {
		result := Top10("a e b e b e c e c e c d d d d")
		require.ElementsMatch(t, result, []string{"e", "d", "c", "b", "a"})
	})

	t.Run("unicode", func(t *testing.T) {
		result := Top10("Луц руЩ ∑ ‱ äöü")
		require.ElementsMatch(t, []string{"луц", "рущ", "äöü"}, result)
	})

	t.Run("tight separation", func(t *testing.T) {
		result := Top10("a.b,c!d$e")
		require.ElementsMatch(t, []string{"a", "b", "c", "d", "e"}, result)
	})

	t.Run("inword separator", func(t *testing.T) {
		result := Top10("a-aa aa-a a-a-a")
		require.ElementsMatch(t, []string{"a-aa", "aa-a", "a-a-a"}, result)
	})

	t.Run("positive", func(t *testing.T) {
		expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
		result := Top10(text)
		require.Subset(t, expected, result)
		require.Len(t, result, 10)
	})
}
