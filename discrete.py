import math
from gmpy2 import mpz, powmod, divm  # type: ignore
import sys


p = mpz(
    13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084171
)
g = mpz(
    11717829880366207009516117596335367088558084999998952205599979459063929499736583746670572176471460312928594829675428279466566527115212748467589894601965568
)
h = mpz(
    3239475104050450443565264378728065788649097520952449527834792452971981976143292558073856937958553180532878928001494706097394108577585732452307673444020333
)


print(f"using [{p=}] | [{g=}] | [{h=}]")

left_hash_table = dict()
found_values = [0, 0]

hash_table_size = int(math.pow(2, 20))
print(f"{hash_table_size=}")

for x1 in range(int(hash_table_size)):

    # h/(g^x1)
    quocient = powmod(g, x1, p)
    left_hand_result = divm(h, quocient, p)
    print(
        f"[{x1}] ({h} / ( {g} ^ {x1} )) -- ( {h} / {quocient} ) -- left hand result = [{left_hand_result}]"
    )

    if left_hand_result not in left_hash_table:
        left_hash_table[left_hand_result] = x1

# sys.exit(0)
# print(f"hash table: {left_hash_table}")

for x0 in range(int(hash_table_size)):

    # (g ^ B) ^ x0
    right_hand_result = powmod(powmod(g, hash_table_size, p), x0, p)
    print(
        f"[{x0}] ( ({g} ^ {hash_table_size}) ^ {x0} ) -- right hand result = [{left_hand_result}]"
    )

    if mpz(right_hand_result) in left_hash_table:
        x1 = left_hash_table[right_hand_result]
        found_values = [x0, x1]
        print(f"found our good value :P --- [x0, x1]={found_values=}")
        break

x = found_values[0] * hash_table_size + found_values[1]

print(f"checking if the values found are good... {found_values=}->{x=}")
pow_check = powmod(g, x, p)
is_good = pow_check == h
print(
    f"result => h = g^x \n\t=> [{h}] = ({g} ^ {x}) \n\t=> [{h}] = [{pow_check}] \n\t=?> {is_good=}"
)
