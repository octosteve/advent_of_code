package day1

import (
	"fmt"
	"strconv"
	strings "strings"
	"testing"
)

func TestSample(t *testing.T) {
	sampleInput := "13399\n13677\n11945\n9861\n6484\n4257\n\n6616\n7349\n7758\n1591\n6068\n9217\n6924\n6766\n\n10040\n9088\n11305\n5867\n10766\n9996\n11092\n\n1320\n4921\n2338\n1351\n3462\n5916\n3124\n1416\n3655\n4886\n1135\n5171\n5020\n5099\n4785\n\n1702\n5083\n3852\n3361\n2505\n3767\n1069\n3564\n3189\n5950\n2250\n2053\n1639\n1430\n4586\n\n4135\n7033\n4649\n3126\n1136\n1435\n3825\n2205\n1259\n5473\n1803\n6406\n\n2466\n30094\n\n3122\n2983\n5988\n4214\n5278\n1974\n7109\n2419\n3777\n8299\n\n10191\n6122\n7298\n7855\n8666\n4777\n6833\n8862\n\n6100\n5332\n1908\n2796\n1818\n4657\n1650\n5560\n8447\n8619\n\n9547\n1354\n1561\n2943\n2547\n9313\n9649\n\n1323\n1200\n2514\n1412\n1381\n2310\n6201\n3257\n1747\n1295\n2507\n5594\n6010\n3422\n\n1995\n3514\n8434\n1625\n2257\n3551\n6881\n\n12256\n1829\n11123\n18829\n\n11790\n1786\n13935\n10088\n3486\n2981\n\n5046\n1306\n4328\n7100\n6000\n6287\n3624\n6521\n3952\n3107\n3846\n2762\n\n21245\n8245\n\n14395\n8107\n4730\n15633\n\n1648\n2062\n4574\n2446\n5471\n3973\n5319\n1370\n4113\n1784\n5596\n3336\n4557\n5690\n4259\n\n9587\n3062\n5611\n10487\n2759\n\n69383\n\n4683\n4950\n1709\n6063\n2338\n4583\n4749\n6431\n4269\n1193\n6012\n5525\n2704\n2462\n\n1627\n7447\n8071\n6915\n7021\n7778\n6861\n4647\n4006\n3953\n\n1866\n9649\n\n10777\n2179\n5406\n1772\n3472\n8632\n\n5609\n7270\n9122\n4723\n8155\n3113\n6348\n6468\n\n5029\n3416\n7531\n8764\n9506\n2104\n7992\n9329\n2530\n\n36312\n11422\n\n5790\n3655\n2172\n6254\n11150\n1624\n1014\n\n6789\n2406\n6639\n4712\n4219\n2880\n5426\n6339\n1460\n1188\n5297\n4006\n2118\n\n3351\n1304\n5559\n3742\n4749\n3648\n5671\n4213\n4047\n1243\n1674\n1943\n4974\n3576\n2816\n\n4949\n6136\n6448\n1652\n2764\n6968\n5423\n3887\n6863\n3628\n4183\n\n2857\n5671\n3486\n4910\n2913\n6441\n5863\n3306\n6473\n2428\n2509\n1188\n3289\n4418\n\n2509\n8521\n2190\n8998\n7303\n5667\n\n3486\n8687\n2976\n5155\n1013\n9599\n8559\n\n9018\n10121\n9235\n5724\n10350\n\n2389\n9061\n8031\n6711\n6046\n3058\n1890\n6047\n7664\n\n1286\n1915\n6597\n\n6012\n6146\n9585\n12335\n3708\n7652\n\n4501\n3470\n3260\n6125\n3718\n3379\n5225\n4673\n5943\n1010\n2057\n3389\n2009\n\n4640\n2437\n4673\n2628\n5333\n7027\n8001\n7678\n4903\n7670\n1898\n\n47670\n\n3884\n23968\n11759\n\n4154\n5214\n2482\n4226\n3115\n2272\n1209\n2525\n2985\n3399\n3140\n4385\n1291\n3911\n\n15947\n12913\n12250\n\n7196\n3574\n6499\n2127\n4295\n6611\n3292\n6542\n1280\n4757\n6430\n7364\n\n36842\n21414\n\n15106\n26273\n\n2751\n1089\n2380\n3818\n1737\n1247\n3897\n1466\n3772\n3890\n4032\n6280\n5202\n3879\n\n4406\n7919\n7174\n7182\n2076\n4281\n9077\n1814\n5768\n\n4183\n5280\n5978\n3859\n1915\n2297\n1478\n3486\n2438\n4829\n2279\n5057\n2625\n5242\n2532\n\n1699\n2463\n5703\n5850\n6899\n2344\n7855\n6715\n6421\n5289\n3702\n\n11196\n4585\n9888\n10474\n10927\n9486\n1240\n\n9145\n35500\n\n9017\n8861\n1486\n5209\n9027\n7222\n8914\n8981\n\n10200\n6668\n9301\n4127\n9401\n5308\n7425\n4684\n\n7333\n10840\n14603\n10690\n\n15280\n12623\n\n2628\n4448\n6702\n5674\n4786\n1955\n5817\n7391\n\n2930\n5391\n1057\n5093\n9249\n5111\n8526\n\n11978\n1172\n12728\n10126\n14320\n\n9024\n10561\n11219\n10949\n9775\n1630\n5280\n\n2785\n6174\n1854\n1905\n6071\n2125\n5350\n1186\n3905\n6166\n2533\n4059\n6731\n\n9855\n3487\n3393\n1922\n3413\n10109\n5879\n1927\n\n24634\n17427\n\n6916\n5891\n4324\n4270\n1627\n3428\n3616\n7358\n\n1720\n7552\n2489\n7386\n2502\n5020\n8134\n6329\n1752\n3651\n\n2666\n4269\n5480\n3812\n1161\n4880\n2090\n1374\n4834\n4310\n\n7381\n7578\n10690\n9863\n\n4168\n6628\n8779\n4512\n9607\n13236\n\n8983\n7035\n\n5321\n6790\n7229\n4435\n4513\n1969\n4220\n6666\n1310\n5702\n2815\n3337\n\n8165\n8740\n12582\n7713\n\n6807\n7574\n2110\n8248\n3745\n8562\n4020\n4169\n1765\n4891\n\n17008\n3541\n13728\n10051\n\n9546\n3634\n11918\n10456\n6942\n8063\n\n8529\n13227\n7794\n1031\n4572\n11405\n\n4153\n6659\n22281\n\n11515\n2370\n\n5548\n1377\n1211\n6916\n1775\n6326\n4643\n5420\n2426\n1870\n6828\n4954\n4020\n\n17677\n4591\n11515\n\n27640\n\n1151\n2701\n3155\n4768\n5921\n1581\n2137\n4733\n5815\n4884\n2760\n2365\n5740\n5709\n4972\n\n8728\n6887\n3823\n2605\n5475\n3860\n6041\n6430\n7326\n1646\n\n4335\n2773\n5938\n2121\n4260\n6369\n6997\n3884\n4772\n1516\n2081\n\n6213\n8805\n7238\n5345\n1192\n4939\n2806\n7313\n\n4753\n6108\n7647\n3770\n1151\n3026\n7908\n4970\n7434\n7037\n2897\n\n5302\n9296\n11044\n7970\n13693\n\n7211\n1990\n3859\n6882\n3344\n4598\n2233\n2877\n2708\n3417\n6936\n2253\n\n6373\n6764\n3884\n7391\n5533\n1137\n2496\n7434\n5125\n3866\n1692\n1925\n\n4935\n6462\n6626\n2070\n1478\n6686\n4058\n2785\n1455\n3868\n\n9644\n4807\n8557\n9675\n3245\n8615\n2544\n\n4673\n6408\n7255\n2955\n8642\n8242\n6319\n4034\n1881\n\n4863\n3401\n1929\n7349\n2101\n3531\n6030\n1829\n2810\n6337\n2740\n3333\n\n2544\n3546\n4223\n3984\n5038\n3259\n3726\n6850\n6201\n2285\n2743\n2805\n\n2831\n5952\n4857\n5209\n3788\n7752\n1897\n5149\n1566\n5627\n1744\n\n13339\n25673\n\n6041\n5270\n3150\n5930\n3880\n7634\n4617\n3736\n1598\n3545\n\n12877\n3897\n16515\n\n8603\n11316\n15009\n14921\n11366\n\n29818\n\n6943\n6084\n3655\n2357\n6786\n4932\n4144\n1856\n6735\n5396\n1969\n3564\n6248\n\n21889\n22636\n7127\n\n5436\n2129\n3427\n4028\n1375\n5428\n3513\n2017\n2241\n2350\n4854\n4469\n6433\n4620\n\n8399\n5055\n7440\n1229\n5012\n8573\n8127\n\n6239\n5302\n1102\n3281\n5757\n5335\n4485\n4455\n4786\n3304\n2231\n1380\n3369\n1580\n\n7307\n4830\n3948\n3827\n1887\n4315\n7452\n4969\n5245\n4550\n7218\n2174\n\n7447\n1081\n4264\n7055\n3586\n5229\n7455\n6934\n6149\n3960\n\n15053\n15764\n7760\n4107\n\n3289\n4370\n9653\n10131\n7939\n7051\n9606\n8231\n\n9300\n8135\n7629\n9080\n3896\n1975\n8696\n1182\n\n3985\n4089\n4921\n5219\n6964\n6433\n6509\n4185\n1005\n1026\n7124\n2661\n\n6051\n14787\n7797\n15612\n10055\n\n4012\n4017\n2304\n2605\n5743\n1403\n1125\n8710\n7720\n8718\n\n6678\n4366\n4875\n6185\n3490\n1883\n3966\n1955\n6868\n2049\n3847\n2604\n5276\n\n3260\n5099\n1168\n4772\n6425\n2118\n4669\n2970\n2545\n2870\n2477\n3117\n5656\n1638\n\n3476\n16997\n3993\n\n1125\n1370\n4944\n8135\n3965\n4389\n9460\n3168\n1910\n\n2953\n1838\n3452\n7390\n5918\n6095\n6666\n7163\n6031\n3723\n\n7718\n1141\n1675\n8446\n5294\n\n5377\n6680\n7310\n6375\n5910\n2649\n9240\n\n7748\n6530\n10511\n12346\n9936\n1147\n\n5383\n2504\n2837\n4482\n2089\n5223\n5901\n2419\n2722\n1275\n5022\n4864\n1186\n3471\n\n1424\n4107\n5138\n6033\n1772\n4181\n3155\n1846\n6531\n6357\n3427\n6064\n\n6840\n12459\n10532\n13685\n7208\n\n10299\n9237\n1345\n5427\n7106\n2016\n1041\n7238\n\n3221\n8456\n3885\n7423\n6113\n4968\n9018\n5032\n6529\n\n33934\n7098\n\n7112\n24553\n\n11885\n11364\n7529\n7103\n8676\n6579\n3688\n\n25617\n10321\n\n4047\n3253\n1844\n2799\n1668\n2965\n4873\n6097\n6295\n6440\n5524\n2244\n5140\n1709\n\n2615\n2931\n3283\n6939\n4277\n2219\n6990\n6804\n6249\n5762\n1461\n1084\n\n32423\n27423\n\n14517\n13142\n10941\n19562\n\n5455\n1050\n5920\n5358\n1310\n4716\n3229\n1832\n4729\n1023\n3679\n2541\n3880\n1570\n3417\n\n16181\n14620\n16387\n14881\n13553\n\n5812\n2876\n4562\n3768\n3935\n5042\n7771\n6679\n7969\n4621\n7996\n\n21785\n\n2879\n3483\n14624\n15533\n\n3101\n1823\n3319\n1134\n5269\n2554\n1437\n6490\n1859\n5525\n1183\n7439\n\n5209\n6916\n5888\n1564\n5739\n6298\n4936\n5691\n3928\n4678\n4320\n1645\n\n2079\n4168\n3031\n2456\n5446\n2338\n1479\n2522\n2715\n1725\n1596\n5845\n5028\n1472\n4262\n\n26049\n35356\n\n7064\n5479\n7564\n6219\n1482\n3370\n2480\n6777\n2460\n4600\n1526\n\n58980\n\n6251\n4189\n4778\n4814\n5817\n1498\n1721\n2071\n1297\n3638\n4299\n6097\n1840\n5523\n\n6232\n2066\n5782\n6549\n2873\n3707\n1461\n3857\n7477\n\n10217\n4464\n7487\n5948\n2641\n6622\n5401\n3488\n\n39956\n\n4731\n4717\n3048\n1349\n1819\n2364\n3993\n4572\n4118\n3673\n2932\n3557\n1818\n4040\n1809\n\n7128\n2747\n10012\n9847\n11104\n\n5461\n8787\n11013\n2435\n7903\n6058\n9698\n\n6006\n2230\n1588\n3453\n1979\n4157\n4378\n2516\n3326\n4336\n5900\n4330\n3758\n1876\n\n6303\n1498\n5453\n3870\n6501\n7204\n7517\n4713\n6341\n2922\n\n7334\n4897\n4402\n10953\n11054\n10268\n9793\n\n4308\n2009\n2882\n6288\n2014\n6973\n5200\n5629\n1899\n5681\n4634\n2955\n\n64804\n\n2857\n3134\n5935\n4192\n3976\n3585\n2199\n4054\n4975\n3372\n1872\n4617\n1449\n2854\n1817\n\n25804\n30702\n\n3722\n15874\n21786\n\n3532\n2397\n6481\n6355\n3696\n6585\n6740\n6653\n1094\n3088\n3531\n\n5927\n6252\n1132\n5256\n2514\n6460\n3890\n5604\n5112\n2472\n2584\n6801\n6849\n\n10325\n4424\n2601\n3731\n1337\n2744\n9597\n7075\n\n8526\n2577\n13053\n18958\n\n4560\n2559\n1153\n4408\n4988\n5067\n5099\n2303\n5451\n5604\n2316\n3957\n4706\n1424\n\n9698\n14099\n2964\n11611\n6559\n\n33997\n17918\n\n3078\n10535\n5480\n8484\n\n9256\n15323\n22521\n\n6174\n1290\n\n2650\n3963\n5043\n2957\n2425\n1608\n5621\n3002\n4820\n7220\n5857\n2229\n\n1787\n9599\n6118\n5536\n10061\n4869\n7507\n10471\n\n6642\n9476\n8827\n1134\n8236\n3198\n9233\n1183\n2111\n\n2519\n13308\n4466\n7787\n11427\n10148\n\n2051\n8362\n3560\n8921\n11983\n10986\n\n17512\n16150\n16427\n4874\n\n1907\n4166\n3506\n3711\n6450\n1079\n6463\n1540\n3003\n1552\n2365\n2677\n6334\n\n5683\n5447\n7822\n1999\n5990\n3230\n6628\n8877\n\n4759\n3993\n1209\n6762\n3003\n6471\n4859\n4760\n1740\n4548\n2611\n4958\n2878\n\n3511\n5142\n2343\n6232\n1655\n7021\n1258\n6838\n6246\n3258\n2833\n5692\n\n14573\n6206\n7184\n\n12569\n12014\n8487\n\n5170\n3993\n2230\n8236\n1047\n2335\n7772\n10473\n\n6108\n6449\n1413\n6964\n3161\n6252\n2130\n2719\n5420\n7277\n\n1583\n8946\n17657\n13041\n\n4422\n2025\n6174\n5449\n3686\n2354\n1209\n4111\n4326\n5606\n2071\n4023\n2194\n1317\n\n4962\n7063\n5352\n4226\n2665\n3288\n6487\n4746\n1582\n3126\n5662\n6864\n\n19194\n18115\n9773\n13571\n\n3436\n3246\n1528\n2100\n5781\n2942\n1985\n3354\n5996\n6033\n4556\n4816\n3445\n5674\n\n7953\n1275\n11734\n3295\n8462\n12028\n10019\n\n1766\n2223\n3111\n3804\n3618\n4227\n2342\n2235\n3909\n2615\n4193\n2296\n\n2843\n1898\n4501\n3483\n6084\n5807\n2422\n5786\n4145\n2572\n5803\n2697\n3218\n2038\n\n11871\n1625\n13110\n12275\n12224\n3092\n\n8174\n8230\n16977\n\n4156\n3989\n4409\n2022\n4922\n5252\n3871\n5072\n4632\n5995\n1056\n4716\n3767\n5414\n1228\n\n1282\n6053\n3736\n2202\n10695\n6392\n1384\n7369\n\n28681\n29246\n\n15208\n21904\n5766\n\n46708\n\n1735\n2394\n6503\n2545\n5642\n1806\n1340\n4804\n5163\n2480\n1842\n5795\n6230\n\n5446\n4579\n4336\n1438\n3998\n2294\n3765\n5777\n4731\n1789\n3080\n5658\n6020\n2542\n3306\n\n11770\n7451\n8351\n7333\n\n25818\n\n23933\n30654\n\n8503\n3608\n18443\n\n1571\n2813\n4847\n5863\n3118\n5888\n5529\n2732\n3260\n6378\n2254\n5451\n3926\n1217\n\n5575\n12012\n8351\n6175\n4427\n5625\n9866\n\n5468\n12325\n21005\n\n5021\n3035\n3699\n1104\n6148\n6925\n6800\n4144\n4039\n5334\n3422\n3349\n2141\n\n18504\n6184\n17915\n17015\n\n2611\n3005\n2768\n2554\n5186\n5515\n4370\n4778\n2033\n4594\n4747\n5050\n6011\n1241\n3687\n\n3572\n2053\n3035\n1019\n4649\n6236\n3675\n1697\n1756\n1384\n6392\n5508\n4879\n2579\n\n17516\n1817\n4564\n\n13428\n11868\n\n2268\n3943\n1945\n4773\n1588\n3347\n2407\n5519\n4959\n2050\n1151\n5578\n3798\n3568\n\n13729\n\n2625\n13582\n7466\n6119\n6529\n8200\n\n5149\n1177\n5792\n7468\n5574\n6245\n5567\n2265\n4336\n7321\n1343\n5623\n\n3722\n8530\n2506\n5780\n5469\n6498\n7749\n2505\n\n22983\n5225\n23416\n\n12172\n\n7005\n2913\n5003\n4441\n8602\n4818\n4073\n7066\n\n2419\n3295\n4377\n5033\n4805\n3757\n1084\n4820\n2652\n4324\n3036\n4980\n1027\n\n4367\n1801\n6250\n7051\n3410\n5676\n3559\n3589\n5903\n4154\n1524\n5197\n\n6557\n4333\n4119\n5285\n4462\n6582\n6615\n5425\n6027\n\n2623\n3519\n4234\n2637\n3220\n2879\n3305\n4309\n2591\n4333\n1030\n2778\n2065\n6071\n2426\n\n8003\n2935\n\n16146\n2907\n25872\n\n2523\n4191\n3738\n5966\n3177\n4487\n4176\n4140\n3472\n1368\n3634\n4151\n3101\n3554\n1097\n\n8706\n7566\n2050\n5536\n1457\n3515\n4303\n9118\n\n2406\n1427\n5485\n14192\n\n6581\n5469\n6933\n5186\n4634\n2457\n8450\n\n6198\n10223\n9913\n1278\n10083\n10253\n8975\n8519\n\n54888\n\n1972\n4402\n4773\n5630\n5896\n6420\n1553\n4579\n1803\n1853\n4043\n5563\n2463\n\n5240\n4200\n2061\n3277\n1346\n2627\n2714\n4837\n4338\n2206\n2051\n5209\n1106\n4793\n3229\n\n6183\n4027\n4577\n4944\n1637\n4339\n2064\n4443\n1125\n2905\n2998\n5512\n6309\n4497\n\n6302\n5477\n3874\n7152\n4289\n1502\n6993\n1968\n8201\n3361\n\n3169\n5005\n1366\n1187\n2960\n4402\n7829\n4494\n4550\n4155\n\n1670\n31707\n\n2778\n6497\n2950\n7022\n3037\n5025\n3920\n1107\n5544\n4234\n1330\n6182\n\n6093\n7193\n1472\n7313\n12034\n\n11527\n11667\n\n8122\n4849\n11876\n6439\n13482\n4171\n\n2510\n5344\n6308\n5470\n3694\n3904\n3060\n1957\n3652\n2225\n1192\n3659\n3671\n2842\n\n5014\n1665\n2556\n1066\n2900\n1324\n2125\n3245\n2785\n3922\n3338\n2431\n2794\n\n6335\n1225\n2642\n6364\n4160\n3309\n\n6834\n12081\n3435\n\n7505\n8207\n3934\n8305\n10764\n11568\n3353\n\n7731\n7976\n3563\n1029\n5288\n1333\n6704\n6018\n1413\n\n1780\n11061\n12941\n5784\n6913\n8216\n\n12014\n5836\n11487\n8744\n10846\n2219\n5146\n\n5868\n3033\n6833\n5479\n5133\n4490\n6107\n5144\n4442\n5649\n2533\n5321\n1877"
	max := 0
	total := 0

	for _, v := range strings.Split(sampleInput, "\n") {
		if v == "" {
			if total > max {
				max = total
			}
			total = 0
		} else {
			i, _ := strconv.Atoi(v)
			total += i
		}
	}

	fmt.Printf("Max is %d\n", max)
	if max != 75622 {
		t.Fail()
	}
}
