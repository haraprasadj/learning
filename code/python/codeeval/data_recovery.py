import sys
if len(sys.argv) > 1:
    test_cases = open(sys.argv[1], 'r')
else:
    test_cases = open('test_cases.txt', 'r')
test_case_list = test_cases.read().strip().split('\n')
for test in test_case_list:
    if test:
        words, numbers = test.split(';')
        word_lst = words.split(' ')
        num_lst = numbers.split(' ')
        num_lst = map(lambda x: int(x), num_lst)
        word_cnt = len(word_lst)
        sum_numbers = sum(num_lst)
        #sum_numbers = reduce(lambda x, y: x+y, numbers)
        missed_num = ((word_cnt * (word_cnt + 1))/2) - sum_numbers
        num_lst.append(missed_num)
        print ' '. join([y for (x, y) in sorted(zip(num_lst, word_lst))])

test_cases.close()