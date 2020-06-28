package xredis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

var (
	Nil = redis.Nil
)

type Client struct {
	cli *redis.Client
}

func NewClient(opt *redis.Options) *Client {
	return &Client{cli: redis.NewClient(opt)}
}

func (c *Client) RawClient() *redis.Client {
	return c.cli
}

//------------------------------------------------------------------------------
func (c *Client) Echo(_ context.Context, message interface{}) *redis.StringCmd {
	return c.cli.Echo(message)
}

func (c *Client) Ping(_ context.Context) *redis.StatusCmd {
	return c.cli.Ping()
}

func (c *Client) Wait(_ context.Context, numSlaves int, timeout time.Duration) *redis.IntCmd {
	return c.cli.Wait(numSlaves, timeout)
}

func (c *Client) Quit(_ context.Context) *redis.StatusCmd {
	panic("not implemented")
}

//------------------------------------------------------------------------------

func (c *Client) Command(_ context.Context) *redis.CommandsInfoCmd {
	return c.cli.Command()
}

func (c *Client) Del(_ context.Context, keys ...string) *redis.IntCmd {
	return c.cli.Del(keys...)
}

func (c *Client) Unlink(keys ...string) *redis.IntCmd {
	return c.cli.Unlink(keys...)
}

func (c *Client) Dump(_ context.Context, key string) *redis.StringCmd {
	return c.cli.Dump(key)
}

func (c *Client) Exists(_ context.Context, keys ...string) *redis.IntCmd {
	return c.cli.Exists(keys...)
}

func (c *Client) Expire(_ context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return c.cli.Expire(key, expiration)
}

func (c *Client) ExpireAt(_ context.Context, key string, tm time.Time) *redis.BoolCmd {
	return c.cli.ExpireAt(key, tm)
}

func (c *Client) Keys(_ context.Context, pattern string) *redis.StringSliceCmd {
	return c.cli.Keys(pattern)
}

func (c *Client) Migrate(_ context.Context, host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd {
	return c.cli.Migrate(host, port, key, db, timeout)
}

func (c *Client) Move(_ context.Context, key string, db int64) *redis.BoolCmd {
	return c.cli.Move(key, db)
}

func (c *Client) ObjectRefCount(_ context.Context, key string) *redis.IntCmd {
	return c.cli.ObjectRefCount(key)
}

func (c *Client) ObjectEncoding(_ context.Context, key string) *redis.StringCmd {
	return c.cli.ObjectEncoding(key)
}

func (c *Client) ObjectIdleTime(_ context.Context, key string) *redis.DurationCmd {
	return c.cli.ObjectIdleTime(key)
}

func (c *Client) Persist(_ context.Context, key string) *redis.BoolCmd {
	return c.cli.Persist(key)
}

func (c *Client) PExpire(_ context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return c.cli.PExpire(key, expiration)
}

func (c *Client) PExpireAt(_ context.Context, key string, tm time.Time) *redis.BoolCmd {
	return c.cli.PExpireAt(key, tm)
}

func (c *Client) PTTL(_ context.Context, key string) *redis.DurationCmd {
	return c.cli.PTTL(key)
}

func (c *Client) RandomKey(_ context.Context) *redis.StringCmd {
	return c.cli.RandomKey()
}

func (c *Client) Rename(_ context.Context, key, newkey string) *redis.StatusCmd {
	return c.cli.Rename(key, newkey)
}

func (c *Client) RenameNX(_ context.Context, key, newkey string) *redis.BoolCmd {
	return c.cli.RenameNX(key, newkey)
}

func (c *Client) Restore(_ context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	return c.cli.Restore(key, ttl, value)
}

func (c *Client) RestoreReplace(_ context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	return c.cli.RestoreReplace(key, ttl, value)
}

func (c *Client) Sort(_ context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	return c.cli.Sort(key, sort)
}

func (c *Client) SortStore(_ context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	return c.cli.SortStore(key, store, sort)
}

func (c *Client) SortInterfaces(_ context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	return c.cli.SortInterfaces(key, sort)
}

func (c *Client) Touch(_ context.Context, keys ...string) *redis.IntCmd {
	return c.cli.Touch(keys...)
}

func (c *Client) TTL(_ context.Context, key string) *redis.DurationCmd {
	return c.cli.TTL(key)
}

func (c *Client) Type(_ context.Context, key string) *redis.StatusCmd {
	return c.cli.Type(key)
}

func (c *Client) Scan(_ context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	return c.cli.Scan(cursor, match, count)
}

func (c *Client) SScan(_ context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return c.cli.SScan(key, cursor, match, count)
}

func (c *Client) HScan(_ context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return c.cli.HScan(key, cursor, match, count)
}

func (c *Client) ZScan(_ context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return c.cli.ZScan(key, cursor, match, count)
}

//------------------------------------------------------------------------------

func (c *Client) Append(_ context.Context, key, value string) *redis.IntCmd {
	return c.cli.Append(key, value)
}

type BitCount struct {
	Start, End int64
}

func (c *Client) BitCount(_ context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	return c.cli.BitCount(key, bitCount)
}

func (c *Client) BitOpAnd(_ context.Context, destKey string, keys ...string) *redis.IntCmd {
	return c.cli.BitOpAnd(destKey, keys...)
}

func (c *Client) BitOpOr(_ context.Context, destKey string, keys ...string) *redis.IntCmd {
	return c.cli.BitOpOr(destKey, keys...)
}

func (c *Client) BitOpXor(_ context.Context, destKey string, keys ...string) *redis.IntCmd {
	return c.cli.BitOpXor(destKey, keys...)
}

func (c *Client) BitOpNot(_ context.Context, destKey string, key string) *redis.IntCmd {
	return c.cli.BitOpNot(destKey, key)
}

func (c *Client) BitPos(_ context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	return c.cli.BitPos(key, bit, pos...)
}

func (c *Client) Decr(_ context.Context, key string) *redis.IntCmd {
	return c.cli.Decr(key)
}

func (c *Client) DecrBy(_ context.Context, key string, decrement int64) *redis.IntCmd {
	return c.cli.DecrBy(key, decrement)
}

// Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (c *Client) Get(_ context.Context, key string) *redis.StringCmd {
	return c.cli.Get(key)
}

func (c *Client) GetBit(_ context.Context, key string, offset int64) *redis.IntCmd {
	return c.cli.GetBit(key, offset)
}

func (c *Client) GetRange(_ context.Context, key string, start, end int64) *redis.StringCmd {
	return c.cli.GetRange(key, start, end)
}

func (c *Client) GetSet(_ context.Context, key string, value interface{}) *redis.StringCmd {
	return c.cli.GetSet(key, value)
}

func (c *Client) Incr(_ context.Context, key string) *redis.IntCmd {
	return c.cli.Incr(key)
}

func (c *Client) IncrBy(_ context.Context, key string, value int64) *redis.IntCmd {
	return c.cli.IncrBy(key, value)
}

func (c *Client) IncrByFloat(_ context.Context, key string, value float64) *redis.FloatCmd {
	return c.cli.IncrByFloat(key, value)
}

func (c *Client) MGet(_ context.Context, keys ...string) *redis.SliceCmd {
	return c.cli.MGet(keys...)
}

func (c *Client) MSet(_ context.Context, pairs ...interface{}) *redis.StatusCmd {
	return c.cli.MSet(pairs...)
}

func (c *Client) MSetNX(_ context.Context, pairs ...interface{}) *redis.BoolCmd {
	return c.cli.MSetNX(pairs...)
}

func (c *Client) Set(_ context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.cli.Set(key, value, expiration)
}

func (c *Client) SetNX(_ context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return c.cli.SetNX(key, value, expiration)
}

func (c *Client) SetXX(_ context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return c.cli.SetXX(key, value, expiration)
}

func (c *Client) SetRange(_ context.Context, key string, offset int64, value string) *redis.IntCmd {
	return c.cli.SetRange(key, offset, value)
}

func (c *Client) StrLen(_ context.Context, key string) *redis.IntCmd {
	return c.cli.StrLen(key)
}

func (c *Client) HDel(_ context.Context, key string, fields ...string) *redis.IntCmd {
	return c.cli.HDel(key, fields...)
}

func (c *Client) HExists(_ context.Context, key, field string) *redis.BoolCmd {
	return c.cli.HExists(key, field)
}

func (c *Client) HGet(_ context.Context, key, field string) *redis.StringCmd {
	return c.cli.HGet(key, field)
}

func (c *Client) HGetAll(_ context.Context, key string) *redis.StringStringMapCmd {
	return c.cli.HGetAll(key)
}

func (c *Client) HIncrBy(_ context.Context, key, field string, incr int64) *redis.IntCmd {
	return c.cli.HIncrBy(key, field, incr)
}

func (c *Client) HIncrByFloat(_ context.Context, key, field string, incr float64) *redis.FloatCmd {
	return c.cli.HIncrByFloat(key, field, incr)
}

func (c *Client) HKeys(_ context.Context, key string) *redis.StringSliceCmd {
	return c.cli.HKeys(key)
}

func (c *Client) HLen(_ context.Context, key string) *redis.IntCmd {
	return c.cli.HLen(key)
}

func (c *Client) HMGet(_ context.Context, key string, fields ...string) *redis.SliceCmd {
	return c.cli.HMGet(key, fields...)
}

func (c *Client) HMSet(_ context.Context, key string, fields map[string]interface{}) *redis.StatusCmd {
	return c.cli.HMSet(key, fields)
}

func (c *Client) HSet(_ context.Context, key, field string, value interface{}) *redis.BoolCmd {
	return c.cli.HSet(key, field, value)
}

func (c *Client) HSetNX(_ context.Context, key, field string, value interface{}) *redis.BoolCmd {
	return c.cli.HSetNX(key, field, value)
}

func (c *Client) HVals(_ context.Context, key string) *redis.StringSliceCmd {
	return c.cli.HVals(key)
}

//------------------------------------------------------------------------------
func (c *Client) BLPop(_ context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return c.cli.BLPop(timeout, keys...)
}

func (c *Client) BRPop(_ context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return c.cli.BRPop(timeout, keys...)
}

func (c *Client) BRPopLPush(_ context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	return c.cli.BRPopLPush(source, destination, timeout)
}

func (c *Client) LIndex(_ context.Context, key string, index int64) *redis.StringCmd {
	return c.cli.LIndex(key, index)
}

func (c *Client) LInsert(_ context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	return c.cli.LInsert(key, op, pivot, value)
}

func (c *Client) LInsertBefore(_ context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return c.cli.LInsertBefore(key, pivot, value)
}

func (c *Client) LInsertAfter(_ context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	return c.cli.LInsertAfter(key, pivot, value)
}

func (c *Client) LLen(_ context.Context, key string) *redis.IntCmd {
	return c.cli.LLen(key)
}

func (c *Client) LPop(_ context.Context, key string) *redis.StringCmd {
	return c.cli.LPop(key)
}

func (c *Client) LPush(_ context.Context, key string, values ...interface{}) *redis.IntCmd {
	return c.cli.LPush(key, values)
}

func (c *Client) LPushX(_ context.Context, key string, value interface{}) *redis.IntCmd {
	return c.cli.LPushX(key, value)
}

func (c *Client) LRange(_ context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return c.cli.LRange(key, start, stop)
}

func (c *Client) LRem(_ context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	return c.cli.LRem(key, count, value)
}

func (c *Client) LSet(_ context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	return c.cli.LSet(key, index, value)
}

func (c *Client) LTrim(_ context.Context, key string, start, stop int64) *redis.StatusCmd {
	return c.cli.LTrim(key, start, stop)
}

func (c *Client) RPop(_ context.Context, key string) *redis.StringCmd {
	return c.cli.RPop(key)
}

func (c *Client) RPopLPush(_ context.Context, source, destination string) *redis.StringCmd {
	return c.cli.RPopLPush(source, destination)
}

func (c *Client) RPush(_ context.Context, key string, values ...interface{}) *redis.IntCmd {
	return c.cli.RPush(key, values)
}

func (c *Client) RPushX(_ context.Context, key string, value interface{}) *redis.IntCmd {
	return c.cli.RPushX(key, value)
}

//------------------------------------------------------------------------------

func (c *Client) SAdd(_ context.Context, key string, members ...interface{}) *redis.IntCmd {
	return c.cli.SAdd(key, members)
}

func (c *Client) SCard(_ context.Context, key string) *redis.IntCmd {
	return c.cli.SCard(key)
}

func (c *Client) SDiff(_ context.Context, keys ...string) *redis.StringSliceCmd {
	return c.cli.SDiff(keys...)
}

func (c *Client) SDiffStore(_ context.Context, destination string, keys ...string) *redis.IntCmd {
	return c.cli.SDiffStore(destination, keys...)
}

func (c *Client) SInter(_ context.Context, keys ...string) *redis.StringSliceCmd {
	return c.cli.SInter(keys...)
}

func (c *Client) SInterStore(_ context.Context, destination string, keys ...string) *redis.IntCmd {
	return c.cli.SInterStore(destination, keys...)
}

func (c *Client) SIsMember(_ context.Context, key string, member interface{}) *redis.BoolCmd {
	return c.cli.SIsMember(key, member)
}

// Redis `SMEMBERS key` command output as a slice
func (c *Client) SMembers(_ context.Context, key string) *redis.StringSliceCmd {
	return c.cli.SMembers(key)
}

// Redis `SMEMBERS key` command output as a map
func (c *Client) SMembersMap(_ context.Context, key string) *redis.StringStructMapCmd {
	return c.cli.SMembersMap(key)
}

func (c *Client) SMove(_ context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	return c.cli.SMove(source, destination, member)
}

// Redis `SPOP key` command.
func (c *Client) SPop(_ context.Context, key string) *redis.StringCmd {
	return c.cli.SPop(key)
}

// Redis `SPOP key count` command.
func (c *Client) SPopN(_ context.Context, key string, count int64) *redis.StringSliceCmd {
	return c.cli.SPopN(key, count)
}

// Redis `SRANDMEMBER key` command.
func (c *Client) SRandMember(_ context.Context, key string) *redis.StringCmd {
	return c.cli.SRandMember(key)
}

// Redis `SRANDMEMBER key count` command.
func (c *Client) SRandMemberN(_ context.Context, key string, count int64) *redis.StringSliceCmd {
	return c.cli.SRandMemberN(key, count)
}

func (c *Client) SRem(_ context.Context, key string, members ...interface{}) *redis.IntCmd {
	return c.cli.SRem(key, members...)
}

func (c *Client) SUnion(_ context.Context, keys ...string) *redis.StringSliceCmd {
	return c.cli.SUnion(keys...)
}

func (c *Client) SUnionStore(_ context.Context, destination string, keys ...string) *redis.IntCmd {
	return c.cli.SUnionStore(destination, keys...)
}

//------------------------------------------------------------------------------
func (c *Client) XAdd(_ context.Context, a *redis.XAddArgs) *redis.StringCmd {
	return c.cli.XAdd(a)
}

func (c *Client) XDel(_ context.Context, stream string, ids ...string) *redis.IntCmd {
	return c.cli.XDel(stream, ids...)
}

func (c *Client) XLen(_ context.Context, stream string) *redis.IntCmd {
	return c.cli.XLen(stream)
}

func (c *Client) XRange(_ context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	return c.cli.XRange(stream, start, stop)
}

func (c *Client) XRangeN(_ context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	return c.cli.XRangeN(stream, start, stop, count)
}

func (c *Client) XRevRange(_ context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	return c.cli.XRevRange(stream, start, stop)
}

func (c *Client) XRevRangeN(_ context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	return c.cli.XRevRangeN(stream, start, stop, count)
}

func (c *Client) XRead(_ context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	return c.cli.XRead(a)
}

func (c *Client) XReadStreams(_ context.Context, streams ...string) *redis.XStreamSliceCmd {
	return c.cli.XReadStreams(streams...)
}

func (c *Client) XGroupCreate(_ context.Context, stream, group, start string) *redis.StatusCmd {
	return c.cli.XGroupCreate(stream, group, start)
}

func (c *Client) XGroupCreateMkStream(_ context.Context, stream, group, start string) *redis.StatusCmd {
	return c.cli.XGroupCreateMkStream(stream, group, start)
}

func (c *Client) XGroupSetID(_ context.Context, stream, group, start string) *redis.StatusCmd {
	return c.cli.XGroupSetID(stream, group, start)
}

func (c *Client) XGroupDestroy(_ context.Context, stream, group string) *redis.IntCmd {
	return c.cli.XGroupDestroy(stream, group)
}

func (c *Client) XGroupDelConsumer(_ context.Context, stream, group, consumer string) *redis.IntCmd {
	return c.cli.XGroupDelConsumer(stream, group, consumer)
}

func (c *Client) XReadGroup(_ context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	return c.cli.XReadGroup(a)
}

func (c *Client) XAck(_ context.Context, stream, group string, ids ...string) *redis.IntCmd {
	return c.cli.XAck(stream, group, ids...)
}

func (c *Client) XPending(_ context.Context, stream, group string) *redis.XPendingCmd {
	return c.cli.XPending(stream, group)
}

func (c *Client) XPendingExt(_ context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	return c.cli.XPendingExt(a)
}

func (c *Client) XClaim(_ context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	return c.cli.XClaim(a)
}

func (c *Client) XClaimJustID(_ context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	return c.cli.XClaimJustID(a)
}

func (c *Client) XTrim(_ context.Context, key string, maxLen int64) *redis.IntCmd {
	return c.cli.XTrim(key, maxLen)
}

func (c *Client) XTrimApprox(_ context.Context, key string, maxLen int64) *redis.IntCmd {
	return c.cli.XTrimApprox(key, maxLen)
}

//------------------------------------------------------------------------------

// ZWithKey represents sorted set member including the name of the key where it was popped.
// Redis `BZPOPMAX key [key ...] timeout` command.
func (c *Client) BZPopMax(_ context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return c.cli.BZPopMax(timeout, keys...)
}

// Redis `BZPOPMIN key [key ...] timeout` command.
func (c *Client) BZPopMin(_ context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	return c.cli.BZPopMin(timeout, keys...)
}

// Redis `ZADD key score member [score member ...]` command.
func (c *Client) ZAdd(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAdd(key, members...)
}

// Redis `ZADD key neon score member [score member ...]` command.
func (c *Client) ZAddNX(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAddNX(key, members...)
}

// Redis `ZADD key XX score member [score member ...]` command.
func (c *Client) ZAddXX(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAddXX(key, members...)
}

// Redis `ZADD key CH score member [score member ...]` command.
func (c *Client) ZAddCh(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAddCh(key, members...)
}

// Redis `ZADD key neon CH score member [score member ...]` command.
func (c *Client) ZAddNXCh(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAddNXCh(key, members...)
}

// Redis `ZADD key XX CH score member [score member ...]` command.
func (c *Client) ZAddXXCh(_ context.Context, key string, members ...redis.Z) *redis.IntCmd {
	return c.cli.ZAddXXCh(key, members...)
}

// Redis `ZADD key INCR score member` command.
func (c *Client) ZIncr(_ context.Context, key string, member redis.Z) *redis.FloatCmd {
	return c.cli.ZIncr(key, member)
}

// Redis `ZADD key neon INCR score member` command.
func (c *Client) ZIncrNX(_ context.Context, key string, member redis.Z) *redis.FloatCmd {
	return c.cli.ZIncrNX(key, member)
}

// Redis `ZADD key XX INCR score member` command.
func (c *Client) ZIncrXX(_ context.Context, key string, member redis.Z) *redis.FloatCmd {
	return c.cli.ZIncrXX(key, member)
}

func (c *Client) ZCard(_ context.Context, key string) *redis.IntCmd {
	return c.cli.ZCard(key)
}

func (c *Client) ZCount(_ context.Context, key, min, max string) *redis.IntCmd {
	return c.cli.ZCount(key, min, max)
}

func (c *Client) ZLexCount(_ context.Context, key, min, max string) *redis.IntCmd {
	return c.cli.ZLexCount(key, min, max)
}

func (c *Client) ZIncrBy(_ context.Context, key string, increment float64, member string) *redis.FloatCmd {
	return c.cli.ZIncrBy(key, increment, member)
}

func (c *Client) ZInterStore(_ context.Context, destination string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return c.cli.ZInterStore(destination, store, keys...)
}

func (c *Client) ZPopMax(_ context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return c.cli.ZPopMax(key, count...)
}

func (c *Client) ZPopMin(_ context.Context, key string, count ...int64) *redis.ZSliceCmd {
	return c.cli.ZPopMin(key, count...)
}

func (c *Client) ZRange(_ context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return c.cli.ZRange(key, start, stop)
}

func (c *Client) ZRangeWithScores(_ context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return c.cli.ZRangeWithScores(key, start, stop)
}

func (c *Client) ZRangeByScore(_ context.Context, key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return c.cli.ZRangeByScore(key, opt)
}

func (c *Client) ZRangeByLex(_ context.Context, key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return c.cli.ZRangeByLex(key, opt)
}

func (c *Client) ZRangeByScoreWithScores(_ context.Context, key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return c.cli.ZRangeByScoreWithScores(key, opt)
}

func (c *Client) ZRank(_ context.Context, key, member string) *redis.IntCmd {
	return c.cli.ZRank(key, member)
}

func (c *Client) ZRem(_ context.Context, key string, members ...interface{}) *redis.IntCmd {
	return c.cli.ZRem(key, members...)
}

func (c *Client) ZRemRangeByRank(_ context.Context, key string, start, stop int64) *redis.IntCmd {
	return c.cli.ZRemRangeByRank(key, start, stop)
}

func (c *Client) ZRemRangeByScore(_ context.Context, key, min, max string) *redis.IntCmd {
	return c.cli.ZRemRangeByScore(key, min, max)
}

func (c *Client) ZRemRangeByLex(_ context.Context, key, min, max string) *redis.IntCmd {
	return c.cli.ZRemRangeByLex(key, min, max)
}

func (c *Client) ZRevRange(_ context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return c.cli.ZRevRange(key, start, stop)
}

func (c *Client) ZRevRangeWithScores(_ context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	return c.cli.ZRevRangeWithScores(key, start, stop)
}

func (c *Client) ZRevRangeByScore(_ context.Context, key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return c.cli.ZRevRangeByScore(key, opt)
}

func (c *Client) ZRevRangeByLex(_ context.Context, key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return c.cli.ZRevRangeByLex(key, opt)
}

func (c *Client) ZRevRangeByScoreWithScores(_ context.Context, key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return c.cli.ZRevRangeByScoreWithScores(key, opt)
}

func (c *Client) ZRevRank(_ context.Context, key, member string) *redis.IntCmd {
	return c.cli.ZRevRank(key, member)
}

func (c *Client) ZScore(_ context.Context, key, member string) *redis.FloatCmd {
	return c.cli.ZScore(key, member)
}

func (c *Client) ZUnionStore(_ context.Context, dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return c.cli.ZUnionStore(dest, store, keys...)
}

//------------------------------------------------------------------------------

func (c *Client) PFAdd(_ context.Context, key string, els ...interface{}) *redis.IntCmd {
	return c.cli.PFAdd(key, els...)
}

func (c *Client) PFCount(_ context.Context, keys ...string) *redis.IntCmd {
	return c.cli.PFCount(keys...)
}

func (c *Client) PFMerge(_ context.Context, dest string, keys ...string) *redis.StatusCmd {
	return c.cli.PFMerge(dest, keys...)
}

func (c *Client) BgRewriteAOF(_ context.Context) *redis.StatusCmd {
	return c.cli.BgRewriteAOF()
}

func (c *Client) BgSave(_ context.Context) *redis.StatusCmd {
	return c.cli.BgSave()
}

func (c *Client) ClientKill(_ context.Context, ipPort string) *redis.StatusCmd {
	return c.cli.ClientKill(ipPort)
}

// ClientKillByFilter is new style syneon, while the ClientKill is old
// CLIENT KILL <option> [value] ... <option> [value]
func (c *Client) ClientKillByFilter(_ context.Context, keys ...string) *redis.IntCmd {
	return c.cli.ClientKillByFilter(keys...)
}

func (c *Client) ClientList(_ context.Context) *redis.StringCmd {
	return c.cli.ClientList()
}

func (c *Client) ClientPause(_ context.Context, dur time.Duration) *redis.BoolCmd {
	return c.cli.ClientPause(dur)
}

func (c *Client) ClientID(_ context.Context) *redis.IntCmd {
	return c.cli.ClientID()
}

func (c *Client) ClientUnblock(_ context.Context, id int64) *redis.IntCmd {
	return c.cli.ClientUnblock(id)
}

func (c *Client) ClientUnblockWithError(_ context.Context, id int64) *redis.IntCmd {
	return c.cli.ClientUnblockWithError(id)
}

// ClientGetName returns the name of the connection.
func (c *Client) ClientGetName(_ context.Context) *redis.StringCmd {
	return c.cli.ClientGetName()
}

func (c *Client) ConfigGet(_ context.Context, parameter string) *redis.SliceCmd {
	return c.cli.ConfigGet(parameter)
}

func (c *Client) ConfigResetStat(_ context.Context) *redis.StatusCmd {
	return c.cli.ConfigResetStat()
}

func (c *Client) ConfigSet(_ context.Context, parameter, value string) *redis.StatusCmd {
	return c.cli.ConfigSet(parameter, value)
}

func (c *Client) ConfigRewrite(_ context.Context) *redis.StatusCmd {
	return c.cli.ConfigRewrite()
}

func (c *Client) DBSize(_ context.Context) *redis.IntCmd {
	return c.cli.DBSize()
}

func (c *Client) FlushAll(_ context.Context) *redis.StatusCmd {
	return c.cli.FlushAll()
}

func (c *Client) FlushAllAsync(_ context.Context) *redis.StatusCmd {
	return c.cli.FlushAllAsync()
}

func (c *Client) FlushDB(_ context.Context) *redis.StatusCmd {
	return c.cli.FlushDB()
}

func (c *Client) FlushDBAsync(_ context.Context) *redis.StatusCmd {
	return c.cli.FlushDBAsync()
}

func (c *Client) Info(_ context.Context, section ...string) *redis.StringCmd {
	return c.cli.Info(section...)
}

func (c *Client) LastSave(_ context.Context) *redis.IntCmd {
	return c.cli.LastSave()
}

func (c *Client) Save(_ context.Context) *redis.StatusCmd {
	return c.cli.Save()
}

func (c *Client) Shutdown(_ context.Context) *redis.StatusCmd {
	return c.cli.Shutdown()
}

func (c *Client) ShutdownSave(_ context.Context) *redis.StatusCmd {
	return c.cli.ShutdownSave()
}

func (c *Client) ShutdownNoSave() *redis.StatusCmd {
	return c.cli.ShutdownNoSave()
}

func (c *Client) SlaveOf(_ context.Context, host, port string) *redis.StatusCmd {
	return c.cli.SlaveOf(host, port)
}

func (c *Client) SlowLog(_ context.Context) {
	panic("not implemented")
}

func (c *Client) Sync(_ context.Context) {
	panic("not implemented")
}

func (c *Client) Time(_ context.Context) *redis.TimeCmd {
	return c.cli.Time()
}

//------------------------------------------------------------------------------

func (c *Client) Eval(_ context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.cli.Eval(script, keys, args...)
}

func (c *Client) EvalSha(_ context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.cli.EvalSha(sha1, keys, args...)
}

func (c *Client) ScriptExists(_ context.Context, hashes ...string) *redis.BoolSliceCmd {
	return c.cli.ScriptExists(hashes...)
}

func (c *Client) ScriptFlush(_ context.Context) *redis.StatusCmd {
	return c.cli.ScriptFlush()
}

func (c *Client) ScriptKill(_ context.Context) *redis.StatusCmd {
	return c.cli.ScriptKill()
}

func (c *Client) ScriptLoad(_ context.Context, script string) *redis.StringCmd {
	return c.cli.ScriptLoad(script)
}

//------------------------------------------------------------------------------

func (c *Client) DebugObject(_ context.Context, key string) *redis.StringCmd {
	return c.cli.DebugObject(key)
}

//------------------------------------------------------------------------------

// Publish posts the message to the channel.
func (c *Client) Publish(_ context.Context, channel string, message interface{}) *redis.IntCmd {
	return c.cli.Publish(channel, message)
}

func (c *Client) PubSubChannels(_ context.Context, pattern string) *redis.StringSliceCmd {
	return c.cli.PubSubChannels(pattern)
}

func (c *Client) PubSubNumSub(_ context.Context, channels ...string) *redis.StringIntMapCmd {
	return c.cli.PubSubNumSub(channels...)
}

func (c *Client) PubSubNumPat(_ context.Context) *redis.IntCmd {
	return c.cli.PubSubNumPat()
}

//------------------------------------------------------------------------------

func (c *Client) ClusterSlots(_ context.Context) *redis.ClusterSlotsCmd {
	return c.cli.ClusterSlots()
}

func (c *Client) ClusterNodes(_ context.Context) *redis.StringCmd {
	return c.cli.ClusterNodes()
}

func (c *Client) ClusterMeet(_ context.Context, host, port string) *redis.StatusCmd {
	return c.cli.ClusterMeet(host, port)
}

func (c *Client) ClusterForget(_ context.Context, nodeID string) *redis.StatusCmd {
	return c.cli.ClusterForget(nodeID)
}

func (c *Client) ClusterReplicate(_ context.Context, nodeID string) *redis.StatusCmd {
	return c.cli.ClusterReplicate(nodeID)
}

func (c *Client) ClusterResetSoft(_ context.Context) *redis.StatusCmd {
	return c.cli.ClusterResetSoft()
}

func (c *Client) ClusterResetHard(_ context.Context) *redis.StatusCmd {
	return c.cli.ClusterResetHard()
}

func (c *Client) ClusterInfo(_ context.Context) *redis.StringCmd {
	return c.cli.ClusterInfo()
}

func (c *Client) ClusterKeySlot(_ context.Context, key string) *redis.IntCmd {
	return c.cli.ClusterKeySlot(key)
}

func (c *Client) ClusterGetKeysInSlot(_ context.Context, slot int, count int) *redis.StringSliceCmd {
	return c.cli.ClusterGetKeysInSlot(slot, count)
}

func (c *Client) ClusterCountFailureReports(_ context.Context, nodeID string) *redis.IntCmd {
	return c.cli.ClusterCountFailureReports(nodeID)
}

func (c *Client) ClusterCountKeysInSlot(_ context.Context, slot int) *redis.IntCmd {
	return c.cli.ClusterCountKeysInSlot(slot)
}

func (c *Client) ClusterDelSlots(_ context.Context, slots ...int) *redis.StatusCmd {
	return c.cli.ClusterDelSlots(slots...)
}

func (c *Client) ClusterDelSlotsRange(_ context.Context, min, max int) *redis.StatusCmd {
	return c.cli.ClusterDelSlotsRange(min, max)
}

func (c *Client) ClusterSaveConfig(_ context.Context) *redis.StatusCmd {
	return c.cli.ClusterSaveConfig()
}

func (c *Client) ClusterSlaves(_ context.Context, nodeID string) *redis.StringSliceCmd {
	return c.cli.ClusterSlaves(nodeID)
}

func (c *Client) ReadOnly(_ context.Context) *redis.StatusCmd {
	return c.cli.ReadOnly()
}

func (c *Client) ReadWrite(_ context.Context) *redis.StatusCmd {
	return c.cli.ReadWrite()
}

func (c *Client) ClusterFailover(_ context.Context) *redis.StatusCmd {
	return c.cli.ClusterFailover()
}

func (c *Client) ClusterAddSlots(_ context.Context, slots ...int) *redis.StatusCmd {
	return c.cli.ClusterAddSlots(slots...)
}

func (c *Client) ClusterAddSlotsRange(_ context.Context, min, max int) *redis.StatusCmd {
	return c.cli.ClusterAddSlotsRange(min, max)
}

//------------------------------------------------------------------------------

func (c *Client) GeoAdd(_ context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return c.cli.GeoAdd(key, geoLocation...)
}

func (c *Client) GeoRadius(_ context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return c.cli.GeoRadius(key, longitude, latitude, query)
}

func (c *Client) GeoRadiusRO(_ context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return c.cli.GeoRadiusRO(key, longitude, latitude, query)
}

func (c *Client) GeoRadiusByMember(_ context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return c.cli.GeoRadiusByMember(key, member, query)
}

func (c *Client) GeoRadiusByMemberRO(_ context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return c.cli.GeoRadiusByMemberRO(key, member, query)
}

func (c *Client) GeoDist(_ context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	return c.cli.GeoDist(key, member1, member2, unit)
}

func (c *Client) GeoHash(_ context.Context, key string, members ...string) *redis.StringSliceCmd {
	return c.cli.GeoHash(key, members...)
}

func (c *Client) GeoPos(_ context.Context, key string, members ...string) *redis.GeoPosCmd {
	return c.cli.GeoPos(key, members...)
}

//------------------------------------------------------------------------------
func (c *Client) MemoryUsage(_ context.Context, key string, samples ...int) *redis.IntCmd {
	return c.cli.MemoryUsage(key, samples...)
}
